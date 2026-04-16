package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
	"zgw/ks/flash_sale/user/configs"
	_ "zgw/ks/flash_sale/user/inits"
	"zgw/ks/flash_sale/user/internal/config"
	"zgw/ks/flash_sale/user/internal/server"
	"zgw/ks/flash_sale/user/internal/svc"
	"zgw/ks/flash_sale/user/pkg"
	"zgw/ks/flash_sale/user/users"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/users.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	pkg.PreheatStock()
	pkg.ScheduledTask()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		users.RegisterUsersServer(grpcServer, server.NewUsersServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
	time.Sleep(time.Second * 2)

}

func DeleteGoods() {
	_, err := configs.Esc.DeleteByQuery("goods").Query(elastic.NewMatchQuery("id", 1)).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("删除商品成功")
}
