package hub

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/contract"
	lerror "github.com/MOSSV2/dimo-sdk-go/lib/error"
	"github.com/MOSSV2/dimo-sdk-go/lib/key"
	"github.com/MOSSV2/dimo-sdk-go/lib/logfs"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/sdk"
	"github.com/gin-gonic/gin"
)

func (s *Server) addUpload(g *gin.RouterGroup) {
	g.Group("/").POST("/upload", func(c *gin.Context) {
		addr := c.PostForm("address")

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}

		if file == nil {
			c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("file is nil")))
			return
		}

		fr, err := file.Open()
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}

		if file.Size == 0 {
			c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("empty file")))
			return
		}
		err = s.logFSWrite(addr, file.Filename, fr)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}

		c.JSON(http.StatusOK, "success")
	})
}

func (s *Server) logFSWrite(addr string, key string, r io.Reader) error {
	var err error
	if addr == "" {
		addr = s.local.String()
	}
	s.Lock()
	fs, ok := s.lfs[addr]
	if !ok {
		fspath := filepath.Join(s.rp.Path(), "logfs")
		fs, err = logfs.New(s.rp.MetaStore(), fspath, addr)
		if err != nil {
			s.Unlock()
			return err
		}
		s.lfs[addr] = fs
	}
	s.Unlock()
	rbytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	err = fs.Put([]byte(key), rbytes)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) logFSRead(addr string, key string, w io.Writer) (int64, error) {
	var err error
	if addr == "" {
		addr = s.local.String()
	}

	s.Lock()
	fs, ok := s.lfs[addr]
	if !ok {
		fspath := filepath.Join(s.rp.Path(), "logfs")
		fs, err = logfs.New(s.rp.MetaStore(), fspath, addr)
		if err != nil {
			s.Unlock()
			return 0, err
		}
		s.lfs[addr] = fs
	}
	s.Unlock()

	lm, err := fs.GetMeta([]byte(key))
	if err != nil {
		return 0, err
	}

	wbytes, err := fs.GetData(lm)
	if err != nil {
		// todo: get from remote
		return 0, err
	}
	n, err := w.Write(wbytes)
	if err != nil {
		return 0, err
	}

	return int64(n), nil
}

func (s *Server) uploadTo() {
	sk := s.rp.Key().Export().PrivateKey
	au, err := key.BuildAuth(sk, []byte("upload"))
	if err != nil {
		panic(err)
	}

	sdk.Login(sdk.ServerURL, au)

	policy := types.Policy{
		N: 6,
		K: 4,
	}

	for {
		time.Sleep(time.Minute)
		err := contract.CheckBalance(au.Addr)
		if err != nil {
			time.Sleep(time.Minute)
			continue
		}

		for _, key := range s.kskeys {
			dsKey := types.NewKey(types.DsLogFS, key)
			val, err := s.rp.MetaStore().Get(dsKey)
			if err != nil || len(val) != 4 {
				continue
			}
			curIndex := binary.BigEndian.Uint32(val)

			next := uint32(0)
			dsKey = types.NewKey(types.DsLogFS, "next", key)
			val, err = s.rp.MetaStore().Get(dsKey)
			if err == nil && len(val) == 4 {
				next = binary.BigEndian.Uint32(val)
			}

			if next >= curIndex {
				continue
			}

			for i := next; i < curIndex; i++ {
				fname := key + "-" + fmt.Sprintf("%d.log", i)
				fp := filepath.Join(s.rp.Path(), "logfs", key, fmt.Sprintf("%d.log", i))

				fr, err := sdk.GetFileReceipt(sdk.ServerURL, au, fname)
				if err == nil {
					er, err := sdk.ListEdge(sdk.ServerURL, au, types.StreamType)
					if err != nil {
						continue
					}
					for _, pn := range fr.Pieces {
						for _, st := range er.Edges {
							pr, err := sdk.GetPieceReceipt(st.ExposeURL, au, pn)
							if err == nil && pr.Serial == 0 {
								err = contract.AddPiece(sk, pr.PieceCore)
								if err != nil {
									continue
								}
							}
						}
					}
					buf := make([]byte, 4)
					binary.BigEndian.PutUint32(buf, i+1)
					s.rp.MetaStore().Put(dsKey, buf)
					continue
				}
				// upload to stream and submit to gateway
				res, streamer, err := sdk.Upload(sdk.ServerURL, au, policy, fp, fname)
				if err != nil {
					break
				}
				pcs, err := sdk.CheckFileFull(res, streamer, fp)
				if err != nil {
					break
				}
				log.Printf("upload %s to %s, sha256: %s\n", fp, streamer, res.Hash)
				log.Printf("submit %s to chain\n", res.Name)
				// submit meta to chain
				for _, pc := range pcs {
					err = contract.AddPiece(sk, pc)
					if err != nil {
						break
					}
				}
				if err != nil {
					break
				}
				buf := make([]byte, 4)
				binary.BigEndian.PutUint32(buf, i+1)
				s.rp.MetaStore().Put(dsKey, buf)
			}
		}
	}
}
