package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"main/app/service/api/internal/config"
	"main/app/service/rpc/login/login"
	"main/app/service/rpc/register/register"
)

type ServiceContext struct {
	Config   config.Config
	Register register.Register
	Login    login.Login
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Register: register.NewRegister(zrpc.MustNewClient(c.Register)),
		Login:    login.NewLogin(zrpc.MustNewClient(c.Login)),
	}
}
