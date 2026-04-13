// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zgw/ks/flash_sale/api/internal/config"
	"zgw/ks/flash_sale/user/usersclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc usersclient.Users
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: usersclient.NewUsers(zrpc.MustNewClient(c.UserRpc)),
	}
}
