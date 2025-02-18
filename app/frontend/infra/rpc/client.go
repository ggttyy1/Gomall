package rpc

import (
	"log"
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartrsevice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/person_infor/personinfor"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/usersevice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	// Client is the RPC client
	UserClient       usersevice.Client
	ProductClient    productcatalogservice.Client
	CartClient       cartrsevice.Client
	CheckoutClient   checkoutservice.Client
	OrderClient      orderservice.Client
	PersonInfoClient personinfor.Client
	once             sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
		initPersonInfoClient()
	})
}

func initPersonInfoClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	if err != nil {
		log.Fatal(err)
	}

	PersonInfoClient, err = personinfor.NewClient("person_infor", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}
func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	if err != nil {
		log.Fatal(err)
	}

	UserClient, err = usersevice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	if err != nil {
		log.Fatal(err)
	}

	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	if err != nil {
		log.Fatal(err)
	}

	CartClient, err = cartrsevice.NewClient("cart", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}

func initCheckoutClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	if err != nil {
		log.Fatal(err)
	}

	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}

func initOrderClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	if err != nil {
		log.Fatal(err)
	}

	OrderClient, err = orderservice.NewClient("order", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}
