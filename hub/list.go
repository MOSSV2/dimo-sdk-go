package hub

import (
	"net/http"
	"strconv"

	lerror "github.com/MOSSV2/dimo-sdk-go/lib/error"
	"github.com/gin-gonic/gin"
)

func (s *Server) addList(g *gin.RouterGroup) {
	g.Group("/").POST("/listAccount", s.listAccountByPost)
	g.Group("/").GET("/listAccount", s.listAccountByGet)

	g.Group("/").POST("/listNeedle", s.listNeedleByPost)
	g.Group("/").GET("/listNeedle", s.listNeedleByGet)

	g.Group("/").POST("/listVolume", s.listVolumeByPost)
	g.Group("/").GET("/listVolume", s.listVolumeByGet)

	g.Group("/").GET("/getAccount", s.getAccountByGet)
	g.Group("/").GET("/getNeedle", s.getNeedleByGet)
	g.Group("/").GET("/getVolume", s.getVolumeByGet)
}

func (s *Server) getAccountByGet(c *gin.Context) {
	owner := c.Query("owner")
	res, err := s.getAccount(owner)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) listAccountByGet(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	length, _ := strconv.Atoi(c.Query("length"))
	if length == 0 {
		length = 32
	}
	res, err := s.listAccount(offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) listAccountByPost(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	length, _ := strconv.Atoi(c.Query("length"))
	if length == 0 {
		length = 32
	}
	res, err := s.listAccount(offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) getNeedleByGet(c *gin.Context) {
	owner := c.Query("owner")
	name := c.Query("name")
	res, err := s.getNeedleDisplay(owner, name)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) listNeedleByGet(c *gin.Context) {
	addr := c.Query("owner")
	offset, _ := strconv.Atoi(c.Query("offset"))
	length, _ := strconv.Atoi(c.Query("length"))
	if length == 0 {
		length = 32
	}
	res, err := s.listNeedleDisplay(addr, offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) listNeedleByPost(c *gin.Context) {
	addr := c.PostForm("owner")
	offset, _ := strconv.Atoi(c.PostForm("offset"))
	length, _ := strconv.Atoi(c.PostForm("length"))
	if length == 0 {
		length = 32
	}
	res, err := s.listNeedleDisplay(addr, offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) listVolumeByGet(c *gin.Context) {
	addr := c.Query("owner")
	offset, _ := strconv.Atoi(c.Query("offset"))
	length, _ := strconv.Atoi(c.Query("length"))
	if length == 0 {
		length = 32
	}
	res, err := s.listVolume(addr, offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) listVolumeByPost(c *gin.Context) {
	addr := c.PostForm("owner")
	offset, _ := strconv.Atoi(c.PostForm("offset"))
	length, _ := strconv.Atoi(c.PostForm("length"))
	if length == 0 {
		length = 32
	}
	res, err := s.listVolume(addr, offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) getVolumeByGet(c *gin.Context) {
	owner := c.Query("owner")
	fid, _ := strconv.Atoi(c.Query("file"))
	res, err := s.getVolume(owner, uint64(fid))
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, res)
}
