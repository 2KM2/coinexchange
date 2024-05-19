package svc

import (
	"common/msdb"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"ucenter-srv/internal/config"
	"ucenter-srv/internal/database"
)

// ServiceContext 配置初始化
type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	Db     *msdb.MsDB
}

func NewServiceContext(c config.Config) *ServiceContext {

	//redis初始化
	//1.Redis服务器
	//2.自定义的Redis命令执行器
	//3.统计缓存的命中率和其他统计数据
	//4.存储分片管理
	//5.配置缓存的选项
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("mscoin"),
		nil,
		func(o *cache.Options) {
		},
	)

	return &ServiceContext{
		Config: c,
		Cache:  redisCache,
		Db:     database.ConnMysql(c.Mysql.DataSource),
	}
}
