package config

import (
	"cron-task/internal/database"
	"cron-task/internal/logic"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Okx        logic.OkxConfig
	Mongo      database.MongoConfig
	Kafka      database.KafkaConfig
	CacheRedis cache.CacheConf
	UCenterRpc zrpc.RpcClientConf
}
