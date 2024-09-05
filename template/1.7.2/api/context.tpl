package svc

import (
	{{.configImport}}
)

type ServiceContext struct {
	Config {{.config}}
    Redis   *redis.Redis
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	return &ServiceContext{
		Config: c,
        Redis:   redis.MustNewRedis(c.Redis),
		{{.middlewareAssignment}}
	}
}
