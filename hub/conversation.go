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

// conversationByPost godoc
//
//	@Summary		get or list conversations
//	@Description	get specific conversation information or list all conversations
//	@Tags			conversation
//	@Accept			json
//	@Produce		json
//	@Param			owner	formData	string	true	"owner/account address"
//	@Param			bucket	formData	string	false	"bucket name"  (empty means list all conversations of owner)
//	@Param			name	formData	string	false	"conversation ID (empty means list all conversation ids)"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/conversation [post]
func (s *Server) conversationByPost(c *gin.Context) {
	ctx := c.Request.Context()
	mn := c.PostForm("name")
	if mn == "" {
		mn = c.PostForm("id")
	}
	addr := c.PostForm("owner")
	bucket := c.PostForm("bucket")
	if mn == "" {
		res, err := s.listConversation(ctx, addr, bucket)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		res, err := s.getConversation(ctx, mn, addr, bucket)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

// conversationByGet godoc
//
//	@Summary		get or list conversations
//	@Description	get specific conversation information or list all conversations
//	@Tags			conversation
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	true	"owner/account address"
//	@Param			bucket	query		string	false	"bucket name" (empty means list all conversations of owner)
//	@Param			name	query		string	false	"conversation ID (empty means list all conversation ids)"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/conversation [get]
func (s *Server) conversationByGet(c *gin.Context) {
	ctx := c.Request.Context()
	mn := c.Query("name")
	if mn == "" {
		mn = c.Query("id")
	}
	addr := c.Query("owner")
	bucket := c.Query("bucket")
	if mn == "" {
		res, err := s.listConversation(ctx, addr, bucket)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		res, err := s.getConversation(ctx, mn, addr, bucket)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
