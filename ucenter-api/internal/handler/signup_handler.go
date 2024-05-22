package handler

import (
	"common"
	"common/tools"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type SignUpHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSignUpHandler(svcCtx *svc.ServiceContext) *SignUpHandler {
	return &SignUpHandler{
		svcCtx: svcCtx,
	}
}

func (sh *SignUpHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req types.Request

	//1.参数解析
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	result := common.NewResult()
	//参数校验
	if req.Captcha == nil {
		httpx.OkJsonCtx(r.Context(), w, result.Deal(nil, errors.New("人机校验不通过")))
		return
	}
	//2.获取IP
	req.Ip = tools.GetRemoteClientIp(r)

	l := logic.NewSignUpLogic(r.Context(), sh.svcCtx)
	resp, err := l.SignUpByPhone(&req)
	if err != nil {
		return
	}

	result.Success("注册成功")
	fmt.Println(resp)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// SendCode 发送验证码-API
func (sh *SignUpHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	var req types.CodeRequest
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	signUpLogic := logic.NewSignUpLogic(r.Context(), sh.svcCtx)
	resp, err := signUpLogic.SendCode(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
