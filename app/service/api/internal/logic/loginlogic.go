package logic

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"main/app/service/rpc/login/app/service/rpc/login"

	"main/app/service/api/internal/svc"
	"main/app/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	res, err := l.svcCtx.Login.Login(l.ctx, &login.LoginReq{
		Uid:      req.Uid,
		Password: req.Password,
	})
	if err != nil {
		g.Log().Errorf(l.ctx, "login failed, err: %v", err)
	}
	return &types.LoginRes{
		Ok:  res.Ok,
		Msg: res.Msg,
	}, nil
}
