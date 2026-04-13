package logic

import (
	"context"
	"errors"
	"log"
	"zgw/ks/flash_sale/user/configs"
	"zgw/ks/flash_sale/user/model"
	"zgw/ks/flash_sale/user/pkg"

	"zgw/ks/flash_sale/user/internal/svc"
	"zgw/ks/flash_sale/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantLoginLogic {
	return &MerchantLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MerchantLoginLogic) MerchantLogin(in *users.MerchantLoginRequest) (*users.MerchantLoginResponse, error) {
	// todo: add your logic here and delete this line
	log.Printf("调用MerchantLogin接口，参数：%v", in)
	var userModel model.User
	if in.Types == 1 {
		if in.Account == "" || in.Password == "" {
			log.Printf("参数不能为空")
			return nil, errors.New("参数不能为空")
		}
		err := userModel.FindUserByAccount(configs.DB, in.Account)
		if err != nil {
			return nil, err
		}
		if pkg.MD5(in.Password) != userModel.Password {
			log.Printf("密码错误")
			return nil, errors.New("密码错误")
		}
	} else {
		if !pkg.IsMobile(in.Mobile) {
			return nil, errors.New("手机号格式错误")
		}
		code, _ := configs.Rdb.Get(l.ctx, in.Mobile).Result()
		if code != in.Code {
			log.Printf("验证码错误")
			return nil, errors.New("验证码错误")
		}
		err := userModel.FindUserByMobile(configs.DB, in.Mobile)
		if err != nil {
			log.Printf("手机号不存在")
			return nil, err
		}
	}

	return &users.MerchantLoginResponse{
		UserId: int64(userModel.ID),
	}, nil
}
