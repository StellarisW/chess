package logic

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"main/app/service/rpc/dao"
	"main/app/service/rpc/register/app/service/rpc/register"
	"main/app/service/rpc/register/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *register.RegisterReq) (*register.RegisterRes, error) {
	cnt, err := g.Model(dao.User.Table()).Where("uid=?", in.Uid).Count()
	if err != nil {
		return &register.RegisterRes{
			Ok:  false,
			Msg: "create user failed, err: " + err.Error(),
		}, err
	}
	if cnt != 0 {
		return &register.RegisterRes{
			Ok:  false,
			Msg: "create user failed, err: user already exist!",
		}, nil
	}
	ePwd := gmd5.MustEncryptString(in.Password)
	_, err = g.Model(dao.User.Table()).Data(g.Map{"uid": in.Uid, "nickname": in.Nickname, "password": ePwd}).Insert()
	if err != nil {
		return &register.RegisterRes{
			Ok:  false,
			Msg: "create user failed, err: " + err.Error(),
		}, err
	}
	return &register.RegisterRes{
		Ok:  true,
		Msg: "create user successfully!",
	}, nil
}
