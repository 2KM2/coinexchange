package logic

import (
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/net/context"
	"grpc-common/ucenter/types/register"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type SignUpLogin struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogin {
	return &SignUpLogin{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpLogin) SignUpByPhone(req *types.Request) (resp *types.Response, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	regReq := &register.RegReq{}
	//将req复制给RPC中定义的类型
	err = copier.Copy(regReq, req)
	if err != nil {
		return nil, err
	}
	//调用RegisterByPhone RPC
	_, err = l.svcCtx.UCRegisterRpc.RegisterByPhone(ctx, regReq)
	if err != nil {
		return nil, err
	}
	return
}

// SendCode 发送验证码-Logic
func (l *SignUpLogin) SendCode(req *types.CodeRequest) (resp *types.CodeResponse, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	_, err = l.svcCtx.UCRegisterRpc.SendCode(ctx, &register.CodeReq{
		Phone:   req.Phone,
		Country: req.Country,
	})
	if err != nil {
		return nil, err
	}
	return
}
