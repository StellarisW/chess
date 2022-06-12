// Code generated by goctl. DO NOT EDIT!
// Source: register.proto

package server

import (
	"context"

	"main/app/service/rpc/register/app/service/rpc/register"
	"main/app/service/rpc/register/internal/logic"
	"main/app/service/rpc/register/internal/svc"
)

type RegisterServer struct {
	svcCtx *svc.ServiceContext
	register.UnimplementedRegisterServer
}

func NewRegisterServer(svcCtx *svc.ServiceContext) *RegisterServer {
	return &RegisterServer{
		svcCtx: svcCtx,
	}
}

func (s *RegisterServer) Register(ctx context.Context, in *register.RegisterReq) (*register.RegisterRes, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}