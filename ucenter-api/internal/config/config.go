package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UCenteredRpc zrpc.RpcClientConf //rpc配置
	JWT          AuthConfig
}
type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
