package admin

import (
	"net/http"

	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/rest/httpx"

	"{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/rpc"
)

// rpc api:
func (m *Service) AdminCall(r *http.Request) (resp *rpc.AdminResp, err error) {
	// TODO: can not use proto req directly, need fix
	var req struct {
		From string `json:"from"`
	}
	if err := httpx.Parse(r, &req); err != nil {
		return nil, err
	}

	resp, err = m.d.Hello.RPC.AdminCall(&rpc.AdminReq{
		From: req.From,
	})
	log.Infof("AdminCall: resp=%+v, err=%v", resp, err)
	return
}
