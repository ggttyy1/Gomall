package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	product "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type UploadProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUploadProductService(Context context.Context, RequestContext *app.RequestContext) *UploadProductService {
	return &UploadProductService{RequestContext: RequestContext, Context: Context}
}

func (h *UploadProductService) Run(req *product.UploadProductReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	file, err := h.RequestContext.FormFile("productImage")
	f, err := file.Open()
	if err != nil {
		log.Fatalf("Error retrieving the file: %v", err)
		return
	}

	// 打印文件名和文件大小
	fmt.Println("File name:", file.Filename)
	fmt.Println("File size:", file.Size)
	path := fmt.Sprintf("./static/image/%s.jpg", file.Filename)
	outFile, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, f)
	if err != nil {
		return nil, err
	}
	_, err = rpc.ProductClient.UploadProduct(h.Context, &rpcproduct.UploadProductRequest{
		Product: &rpcproduct.UpProduct{
			Name:            req.Name,
			Price:           req.Price,
			Description:     req.Description,
			Picture:         path,
			ProductCategory: req.ProductCategory,
		},
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"Title": "home",
	}, nil
}
