package rpc

import (
	"fmt"
	"log"
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/cart/conf"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	// Client is the RPC client
	ProductClient productcatalogservice.Client
	once          sync.Once
	err           error
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

var (
	CurrentServiceName = conf.GetConf().Kitex.Service
	RegistryAddress    = conf.GetConf().Registry.RegistryAddress[0]
)

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(&clientsuite.CommonClientSuite{
			CurrentServiceName: CurrentServiceName,
			RegistryAddress:    RegistryAddress,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ProductClient init success")
}

func TestInit() {
	once.Do(func() {
		TestinitProductClient()
	})
}
func TestinitProductClient() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}

	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ProductClient init success")
}
