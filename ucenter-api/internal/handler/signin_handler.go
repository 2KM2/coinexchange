package handler

import (
	"common"
	"common/tools"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type SignInHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSignInHandler(svcCtx *svc.ServiceContext) *SignInHandler {
	return &SignInHandler{
		svcCtx: svcCtx,
	}
}

// SignIn 登录-API
func (sh *SignInHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req types.LoginReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	newResult := common.NewResult()
	if req.Captcha == nil {
		httpx.OkJsonCtx(r.Context(), w, newResult.Deal(nil, errors.New("人机校验不通过")))
		return
	}
	//2.获取IP
	req.Ip = tools.GetRemoteClientIp(r)
	//3.登录
	l := logic.NewSignInLogic(r.Context(), sh.svcCtx)
	resp, err := l.SignInByPhone(&req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (sh *SignInHandler) CheckLogin(w http.ResponseWriter, r *http.Request) {
	logx.Info("CheckLogin")
	newResult := common.NewResult()
	token := r.Header.Get("x-auth-token")
	l := logic.NewSignInLogic(r.Context(), sh.svcCtx)
	//data : true or false
	resp, err := l.CheckLogin(token)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
