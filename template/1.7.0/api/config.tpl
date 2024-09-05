package config

import {{.authImport}}

type Config struct {
	rest.RestConf
	Redis   redis.RedisConf
	{{.auth}}
	{{.jwtTrans}}
}
