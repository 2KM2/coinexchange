package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      MysqlConfig
	CacheRedis cache.CacheConf
	Captcha    CaptchaConf
	//Consul     consul.Conf
	JWT AuthConfig
}
type MysqlConfig struct {
	DataSource string
}
type CaptchaConf struct {
	Vid string
	Key string
}
type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
