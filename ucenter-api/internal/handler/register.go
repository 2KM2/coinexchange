package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"ucenter-api/internal/svc"
)

func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	//注册路由组
	logx.Info("NewSignUpHandler")
	signup := NewSignUpHandler(serverCtx)
	signUpGroups := r.Group()
	signUpGroups.Post("/uc/register/phone", signup.SignUp)
	signUpGroups.Post("/uc/mobile/code", signup.SendCode)

	//登录路由
	logx.Info("NewSignInHandler")
	signIn := NewSignInHandler(serverCtx)
	signInGroups := r.Group()
	signInGroups.Post("/uc/login", signIn.SignIn)
	signInGroups.Post("/uc/check/login", signIn.CheckLogin)
}
