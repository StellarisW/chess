package logic

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"main/app/service/rpc/dao"
	"main/app/service/rpc/login/app/service/rpc/login"
	"main/app/service/rpc/login/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginRes, error) {
	cnt, err := g.Model(dao.User.Table()).Where("uid=?", in.Uid).Count()
	if err != nil {
		return &login.LoginRes{
			Ok:  false,
			Msg: "login failed, err: " + err.Error(),
		}, err
	}
	if cnt == 0 {
		return &login.LoginRes{
			Ok:  false,
			Msg: "login failed, err: uid doesn't exist!",
		}, nil
	}
	ePwd := gmd5.MustEncryptString(in.Password)
	r, err := g.Model(dao.User.Table()).Where("uid=? and password=?", in.Uid, ePwd).One()
	if err != nil {
		return &login.LoginRes{
			Ok:  false,
			Msg: "login failed, err: " + err.Error(),
		}, err
	}
	if len(r) == 0 {
		return &login.LoginRes{
			Ok:  false,
			Msg: "login failed, err: wrong password",
		}, nil
	}
	return &login.LoginRes{
		Ok:  true,
		Msg: "login successfully!",
	}, nil
}
