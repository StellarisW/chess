package logic

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"main/app/service/rpc/register/register"

	"main/app/service/api/internal/svc"
	"main/app/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	res, err := l.svcCtx.Register.Register(l.ctx, &register.RegisterReq{
		Uid:      req.Uid,
		Nickname: req.Nickname,
		Password: req.Password,
	})
	if err != nil {
		g.Log().Errorf(l.ctx, "create user failed, err: %v", err)
	}
	return &types.RegisterRes{
		Ok:  res.Ok,
		Msg: res.Msg,
	}, nil
}
