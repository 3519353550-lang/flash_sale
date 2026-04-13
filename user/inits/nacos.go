package inits

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"strings"
	"zgw/ks/flash_sale/user/configs"
)

func InitNacos() {
	viper.SetConfigFile("../user/configs/dev.yaml")
	viper.ReadInConfig()
	var Nacos configs.Nacos
	viper.UnmarshalKey("NaCos", &Nacos)
	fmt.Println("Nacos:", Nacos)

	// Nacos服务器地址
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: Nacos.Host,
			Port:   uint64(Nacos.Port),
		},
	}
	// 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         Nacos.Namespace, // 如果不需要命名空间，可以留空
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            Nacos.Username,
		Password:            Nacos.Password,
	}

	// 创建配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	dataId := Nacos.DataId
	group := Nacos.Group
	config, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("configs", config)

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "dataId",
		Group:  "group",
		OnChange: func(namespace, group, dataId, data string) {
			//fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		panic(err)
	}

	viper.Reset()
	viper.SetConfigType("json")
	viper.ReadConfig(strings.NewReader(config))
	viper.Unmarshal(&configs.Conf)
	fmt.Println("configs", configs.Conf)
}
