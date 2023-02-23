package svc

import (
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/greet/internal/config"
	"go-zero-demo/greet/internal/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config          config.Config
	Db              *gorm.DB
	Rdb             *redis.Client
	OperationRecord rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	//mysql 上下文加载 数据库配置
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:               c.Database.DSN,
		DefaultStringSize: 171, //数据库varchar类型的默认值
	}), &gorm.Config{
		SkipDefaultTransaction: false, //启用事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   //表前缀
			SingularTable: true, //使用单数表名
		},
		DryRun:                                   false,                               //禁止SQL空跑
		DisableForeignKeyConstraintWhenMigrating: true,                                //创建逻辑外键
		Logger:                                   logger.Default.LogMode(logger.Info), //输出 SQL语句
	})

	//redis 上下文加载 数据库配置
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host + ":" + c.Redis.Port,
		Password: c.Redis.Password,
		DB:       c.Redis.Db,
	})

	return &ServiceContext{
		Config:          c,
		Db:              db,
		Rdb:             rdb,
		OperationRecord: middleware.NewOperationRecord().Handle,
	}
}
