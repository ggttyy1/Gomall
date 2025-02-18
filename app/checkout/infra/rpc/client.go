package rpc

import (
	"fmt"
	"log"
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/cart/conf"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartrsevice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor/personinfor"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	// Client is the RPC client
	ProductClient      productcatalogservice.Client
	CartClient         cartrsevice.Client
	PaymentClient      paymentservice.Client
	OrderClient        orderservice.Client
	PersonInfoClient   personinfor.Client
	once               sync.Once
	CurrentServiceName = conf.GetConf().Kitex.Service
	RegistryAddress    = conf.GetConf().Registry.RegistryAddress[0]
	err                error
)

func Init() {
	once.Do(func() {
		initProductClient()
		initCartClient()
		initPaymentClient()
		initOrderClient()
		initPersonInfoClient()
	})
}

func initPersonInfoClient() {
	opts := []client.Option{
		client.WithSuite(&clientsuite.CommonClientSuite{
			CurrentServiceName: CurrentServiceName,
			RegistryAddress:    RegistryAddress,
		}),
	}

	PersonInfoClient, err = personinfor.NewClient("person_infor", opts...)
	if err != nil {
		log.Fatal(err)
	}
}

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
}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(&clientsuite.CommonClientSuite{
			CurrentServiceName: CurrentServiceName,
			RegistryAddress:    RegistryAddress,
		}),
	}

	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		log.Fatal(err)
	}
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(&clientsuite.CommonClientSuite{
			CurrentServiceName: CurrentServiceName,
			RegistryAddress:    RegistryAddress,
		}),
	}

	CartClient, err = cartrsevice.NewClient("cart", opts...)
	if err != nil {
		log.Fatal(err)
	}
}

func initPaymentClient() {
	opts := []client.Option{
		client.WithSuite(&clientsuite.CommonClientSuite{
			CurrentServiceName: CurrentServiceName,
			RegistryAddress:    RegistryAddress,
		}),
	}

	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		log.Fatal(err)
	}
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
