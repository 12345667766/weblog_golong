package views

import (
	"ms-go-blog/common"
	"ms-go-blog/config"
	"ms-go-blog/context"
	"net/http"
)

func (*HTMLApi) LoginNew(ctx *context.MsContext) {
	login := common.Template.Login

	login.WriteData(ctx.W, config.Cfg.Viewer)
}
func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)
}

func (*HTMLApi) Register(w http.ResponseWriter, r *http.Request) {
	register := common.Template.Register

	register.WriteData(w, config.Cfg.Viewer)
}
