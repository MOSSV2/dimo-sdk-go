package hub

import (
	"net/http"

	lerror "github.com/MOSSV2/dimo-sdk-go/lib/error"
	"github.com/gin-gonic/gin"
)

func (s *Server) addConversation(g *gin.RouterGroup) {
	g.Group("/").POST("/conversation", s.conversationByPost)
	g.Group("/").GET("/conversation", s.conversationByGet)
}

func (s *Server) conversationByPost(c *gin.Context) {
	ctx := c.Request.Context()
	mn := c.PostForm("name")
	if mn == "" {
		mn = c.PostForm("id")
	}
	addr := c.PostForm("owner")

	if mn == "" {
		res, err := s.listConversation(ctx, addr)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		res, err := s.getConversation(ctx, mn, addr)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func (s *Server) conversationByGet(c *gin.Context) {
	ctx := c.Request.Context()
	mn := c.Query("name")
	if mn == "" {
		mn = c.Query("id")
	}
	addr := c.Query("owner")

	if mn == "" {
		res, err := s.listConversation(ctx, addr)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		res, err := s.getConversation(ctx, mn, addr)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
