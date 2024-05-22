package logic

import (
	"common/tools"
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/ucenter/types/login"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignInByPhone(req *types.LoginReq) (resp *types.LoginRes, err error) {
	logx.Info("[SignInLogic] SignInByPhone", req)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	signInReq := &login.LoginReq{}
	if err := copier.Copy(signInReq, req); err != nil {
		return nil, err
	}
	//调用登录的RPC
	signInResp, err := l.svcCtx.UCLoginRpc.Login(ctx, signInReq)
	if err != nil {
		return nil, err
	}
	resp = &types.LoginRes{}
	if err := copier.Copy(resp, signInResp); err != nil {
		return nil, err
	}
	return
}

func (l *SignInLogic) CheckLogin(token string) (bool, error) {
	_, err := tools.ParseToken(token, l.svcCtx.Config.JWT.AccessSecret)
	if err != nil {
		logx.Error(err)
		return false, nil
	}
	return true, nil
}
