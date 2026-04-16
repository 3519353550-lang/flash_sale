package logic

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"time"
	"zgw/ks/flash_sale/user/configs"
	"zgw/ks/flash_sale/user/pkg"

	"zgw/ks/flash_sale/user/internal/svc"
	"zgw/ks/flash_sale/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *users.SendSmsRequest) (*users.SendSmsResponse, error) {
	// todo: add your logic here and delete this line
	//if !pkg.IsMobile(in.Mobile) {
	//	return nil, errors.New("手机号格式错误")
	//}
	data, _ := configs.Rdb.Get(l.ctx, "sms:"+in.Mobile).Int()
	if data >= 3 {
		return nil, errors.New("短信发送频率过快")
	}

	code := rand.Intn(900000) + 100000

	sms, err := pkg.SendSms(in.Mobile, strconv.Itoa(code))
	if err != nil {

		return nil, err
	}
	if sms.Code != 2 {
		return nil, errors.New(sms.Msg)
	}

	configs.Rdb.Set(l.ctx, in.Mobile, code, 5*time.Minute)
	configs.Rdb.Set(l.ctx, "sms:"+in.Mobile, 3, 1*time.Minute)

	return &users.SendSmsResponse{Success: true}, nil
}
