package database

import (
	"common/msdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnMysql(dsn string) *msdb.MsDB {
	var err error
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败,error=" + err.Error())
	}
	db, _ := _db.DB()

	//配置连接池
	db.SetMaxOpenConns(200) // 从较低的数字开始，根据需要逐步增加
	db.SetMaxIdleConns(50)  // 保持一个合理数量的空闲连接
	return &msdb.MsDB{
		Conn: _db,
	}
}
