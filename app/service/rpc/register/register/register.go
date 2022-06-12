// Code generated by goctl. DO NOT EDIT!
// Source: register.proto

package register

import (
	"context"

	"main/app/service/rpc/register/app/service/rpc/register"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	RegisterReq = register.RegisterReq
	RegisterRes = register.RegisterRes

	Register interface {
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
	}

	defaultRegister struct {
		cli zrpc.Client
	}
)

func NewRegister(cli zrpc.Client) Register {
	return &defaultRegister{
		cli: cli,
	}
}

func (m *defaultRegister) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error) {
	client := register.NewRegisterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}
