gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module github.com/cloudwego/biz-demo/gomall/app/frontend --idl ../../idl/frontend/auth_page.proto

gen-user_client:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/user.proto

gen-user-server:
	@cd app/user && cwgo server --type RPC --service user --module github.com/cloudwego/biz-demo/gomall/app/user --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl --idl ../../idl/user.proto

gen-product_client:
	@cd rpc_gen && cwgo client --type RPC --service product --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/product.proto

gen-product-server:
	@cd app/product && cwgo server --type RPC --service product --module github.com/cloudwego/biz-demo/gomall/app/product --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl --idl ../../idl/product.proto

gen-frontend-product:
	cwgo server -I ../../idl --type HTTP --service frontend --module github.com/cloudwego/biz-demo/gomall/app/frontend --idl ../../idl/frontend/product_page.proto
	cwgo server -I ../../idl --type HTTP --service frontend --module github.com/cloudwego/biz-demo/gomall/app/frontend --idl ../../idl/frontend/category_page.proto

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --service cart --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module github.com/cloudwego/biz-demo/gomall/app/cart --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl --idl ../../idl/cart.proto

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --service payment --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --service payment --module github.com/cloudwego/biz-demo/gomall/app/payment --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl --idl ../../idl/payment.proto

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --service checkout --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --service checkout --module github.com/cloudwego/biz-demo/gomall/app/checkout --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl --idl ../../idl/checkout.proto

.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --service order --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module github.com/cloudwego/biz-demo/gomall/app/order --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl --idl ../../idl/order.proto

.PHONY: gen-person_infor
gen-person_infor:
	@cd rpc_gen && cwgo client --type RPC --service person_infor --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/person_infor.proto
	@cd app/person_infor && cwgo server --type RPC --service person_infor --module github.com/cloudwego/biz-demo/gomall/app/person_infor --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl --idl ../../idl/person_infor.proto
.PHONY: gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --service email --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC --service email --module github.com/cloudwego/biz-demo/gomall/app/email --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl --idl ../../idl/email.proto


.PHONY: gen-person_infor
gen-person_infor:
	@cd rpc_gen && cwgo client --type RPC --service person_infor --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/person_infor.proto
	@cd app/person_infor && cwgo server --type RPC --service person_infor --module github.com/cloudwego/biz-demo/gomall/app/person_infor --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl --idl ../../idl/person_infor.proto
