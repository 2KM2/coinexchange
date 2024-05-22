package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/ucenter/ucclient"
	"ucenter-api/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UCRegisterRpc ucclient.Register
	Cache         cache.Cache
	UCLoginRpc    ucclient.Login //登录的Client-RPC
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UCRegisterRpc: ucclient.NewRegister(zrpc.MustNewClient(c.UCenteredRpc)),
		UCLoginRpc:    ucclient.NewLogin(zrpc.MustNewClient(c.UCenteredRpc)),
	}
}
