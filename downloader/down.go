package downloader

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/MOSSV2/dimo-sdk-go/build"
	lerror "github.com/MOSSV2/dimo-sdk-go/lib/error"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/sdk"
	"github.com/gin-gonic/gin"
)

func (s Server) addDownload(g *gin.RouterGroup) {
	g.Group("/").POST("/download", s.downloadByPOST)
	g.Group("/").GET("/download", s.downloadByGET)
}

func (s Server) downloadByGET(c *gin.Context) {
	ctx := c.Request.Context()
	mn := c.Query("name")

	head := fmt.Sprintf("attachment; filename=\"%s\"", mn)
	extraHeaders := map[string]string{
		"Content-Disposition": head,
	}

	var w bytes.Buffer
	res, err := s.ps.GetReplica(ctx, mn, &w, types.Options{})
	if err != nil {
		fr, err := sdk.GetFileReceipt(build.ServerURL, s.auth, mn)
		if err == nil {
			err = sdk.Download(build.ServerURL, s.auth, mn, s.ps, &w)
			if err != nil {
				c.JSON(599, lerror.ToAPIError("file", err))
				return
			}
			c.DataFromReader(http.StatusOK, fr.Size, "text/plain", &w, extraHeaders)
			return
		} else {
			pr, err := sdk.GetPieceReceipt(build.ServerURL, s.auth, mn)
			if err == nil {
				err = sdk.DownloadPieceAndSave(build.ServerURL, s.auth, mn, s.ps)
				if err != nil {
					c.JSON(599, lerror.ToAPIError("file", err))
					return
				}

				_, err := s.ps.GetPiece(context.TODO(), mn, &w, types.Options{})
				if err != nil {
					c.JSON(599, lerror.ToAPIError("file", err))
					return
				}
				c.DataFromReader(http.StatusOK, pr.Size, "text/plain", &w, extraHeaders)
				return
			}
		}
	}

	c.DataFromReader(http.StatusOK, res.Size, "text/plain", &w, extraHeaders)
}

func (s Server) downloadByPOST(c *gin.Context) {

}
