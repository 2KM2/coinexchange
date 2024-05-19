package handler

import (
	"ucenter-api/internal/svc"
)

func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	//注册路由组
	signup := NewSignUpHandler(serverCtx)
	signUpGroups := r.Group()
	signUpGroups.Post("/uc/register/phone", signup.SignUp)
	signUpGroups.Post("/uc/mobile/code", signup.SendCode)
}
