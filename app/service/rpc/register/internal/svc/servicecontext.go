package svc

import (
	"main/app/service/rpc/register/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//Model:  model.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
