package hub

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/log"
	"github.com/MOSSV2/dimo-sdk-go/lib/logfs"
	"github.com/MOSSV2/dimo-sdk-go/lib/piece"
	"github.com/MOSSV2/dimo-sdk-go/lib/repo"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
	"github.com/MOSSV2/dimo-sdk-go/sdk"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"gorm.io/gorm"
)

var logger = log.Logger("hub")

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

type Server struct {
	Router *gin.Engine

	typ string

	rp repo.Repo

	gdb *gorm.DB

	ps types.IPieceStore

	sync.RWMutex
	fscnt uint32
	lfs   map[string]*logfs.LogFS

	local common.Address

	auth types.Auth
}

func NewServer(rp repo.Repo) (*http.Server, error) {
	log.SetLogLevel("DEBUG")

	gin.SetMode(gin.ReleaseMode)

	localAddr := rp.Key().Address()

	logger.Infof("hub %s starting...", localAddr)

	router := gin.Default()

	auth, err := rp.Key().BuildAuth([]byte("hub"))
	if err != nil {
		return nil, err
	}

	s := &Server{
		Router: router,

		typ:   types.HubType,
		local: localAddr,
		rp:    rp,
		ps:    piece.New(rp.MetaStore(), rp.DataStore()),
		auth:  auth,

		lfs: make(map[string]*logfs.LogFS),
	}

	err = s.register()
	if err != nil {
		return nil, err
	}

	s.loadGORM()

	s.load()
	go s.uploadTo()

	s.registRoute()

	srv := &http.Server{
		Addr:    rp.Config().API.Endpoint,
		Handler: s.Router,
	}

	return srv, nil
}

func (s *Server) registRoute() {
	swaghost := s.rp.Config().API.Expose
	if swaghost != "" {
		swaghost = strings.TrimPrefix(swaghost, "http://")
		docs.SwaggerInfo.Host = swaghost
	}

	s.Router.Use(Cors())

	s.Router.Use(ginzap.Ginzap(log.Logger("gin").Desugar(), time.RFC3339, true))

	s.Router.Use(static.Serve("/", static.LocalFile("assets", true)))

	s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r := s.Router.Group("/api")

	s.addInfo(r)
	s.addDownload(r)
	s.addUpload(r)
	s.addList(r)
}

func login(url string, auth types.Auth) {
	for {
		sdk.Login(url, auth)
		time.Sleep(time.Hour)
	}
}

func (s *Server) register() error {
	auth, err := s.rp.Key().BuildAuth([]byte("register"))
	if err != nil {
		return err
	}

	go login(s.rp.Config().Remote.URL, auth)

	mm := types.EdgeMeta{
		Type:      s.typ,
		Name:      auth.Addr,
		PublicKey: s.rp.Key().Public(),
		ExposeURL: s.rp.Config().API.Expose,
		Hardware:  utils.GetHardwareInfo(),
	}

	err = sdk.RegisterEdge(s.rp.Config().Remote.URL, auth, mm)
	if err != nil {
		logger.Debug("register hub fail:", err)
		return err
	}
	return nil
}

func (s *Server) addInfo(g *gin.RouterGroup) {
	g.Group("/").GET("/info", func(c *gin.Context) {
		res := types.EdgeReceipt{
			EdgeMeta: types.EdgeMeta{
				Type: s.typ,
				Name: s.local,
			},
		}

		c.JSON(http.StatusOK, res)
	})
}
