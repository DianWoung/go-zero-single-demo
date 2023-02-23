package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	//config.yaml格式 会转换成 struct格式
	Database struct {
		DSN             string
		MaxOpenConn     int
		MaxIdleConn     int
		ConnMaxLifetime int
		ConnMaxIdleTime int
	}

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Redis struct {
		Host     string
		Password string
		Port     string
		Db       int
	}
}
