package svc

import {{.imports}}

type ServiceContext struct {
	Config config.Config
    Redis:   redis.MustNewRedis(c.Redis),
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:c,
        Redis:   redis.MustNewRedis(c.Redis),
        // BookModel: models.NewBooksModel(sqlx.NewMysql(c.Mysql.DataSource), c.Cache),
	}
}
