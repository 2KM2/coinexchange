package ucclient

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"grpc-common/ucenter/types/login"
)

type (
	loginReq = login.LoginReq
	loginRes = login.LoginRes

	Login interface {
		Login(ctx context.Context, in *loginReq, opts ...grpc.CallOption) (*loginRes, error)
	}

	defaultLogin struct {
		cli zrpc.Client
	}
)

func NewLogin(cli zrpc.Client) Login {
	return &defaultLogin{
		cli: cli,
	}
}
func (m *defaultLogin) Login(ctx context.Context, in *loginReq, opts ...grpc.CallOption) (*loginRes, error) {
	client := login.NewLoginClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}