package hub

import (
	"net/http"
	"strconv"

	lerror "github.com/MOSSV2/dimo-sdk-go/lib/error"
	"github.com/gin-gonic/gin"
)

func (s *Server) addList(g *gin.RouterGroup) {
	g.Group("/").POST("/listAccount", s.listAccount)
	g.Group("/").GET("/listAccount", s.listAccount)

	g.Group("/").POST("/listFile", s.listFileByPost)
	g.Group("/").GET("/listFile", s.listFile)
}

func (s *Server) listAccount(c *gin.Context) {
	res, err := s.listAccountInGORM()
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) listFile(c *gin.Context) {
	addr := c.Query("owner")
	offset, _ := strconv.Atoi(c.Query("offset"))
	length, _ := strconv.Atoi(c.Query("length"))
	if length == 0 {
		length = 32
	}
	res, err := s.listNeedle(addr, offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) listFileByPost(c *gin.Context) {
	addr := c.PostForm("owner")
	offset, _ := strconv.Atoi(c.PostForm("offset"))
	length, _ := strconv.Atoi(c.PostForm("length"))
	if length == 0 {
		length = 32
	}
	res, err := s.listNeedle(addr, offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}
