package tools

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
	"github.com/pkg/browser"
)

// 获取添加 todo 工具
// 使用 utils.NewTool 创建工具

type OpenWebParams struct {
	URL string `json:"url"` // 用于打开的 URL 地址
}

// 定义打开 CloudWeGo 网站的工具
func awakeClodeWeGo() tool.InvokableTool {
	info := &schema.ToolInfo{
		Name: "open_cloudwego",
		Desc: "Open CloudWeGo website",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"url": {
				Desc:     "The URL of the website to open",
				Type:     schema.String, // 网站的URL参数是字符串
				Required: true,
			},
		}),
	}

	// 使用 utils.NewTool 来创建工具，绑定到 OpenWeb 函数
	return utils.NewTool(info, OpenWeb)
}

// 定义处理打开网站逻辑的函数
func OpenWeb(_ context.Context, params *OpenWebParams) (string, error) {
	Infof("Opening website with URL: %+v", params)

	// 假设你使用 web 浏览器或某些工具打开网址
	// 这里我们只是打印参数内容，实际上你可能需要调用外部库来实现打开网页的功能
	url := params.URL // 假设 Content 是传入的 URL 参数
	if url == "" {
		return `{"msg": "URL is required"}`, fmt.Errorf("URL is required")
	}

	// 在这里你可以实现调用实际打开网址的代码，假设是用浏览器打开
	// 例如：webbrowser.Open(url) 或其他方式打开网页
	err := browser.OpenURL("https://openai.com/chatgpt/overview/")
	if err != nil {
		return "", err
	}
	// 返回成功消息
	return `{"msg": "Open website success"}`, nil
}

func createHomeJumpTool() tool.InvokableTool {
	info := &schema.ToolInfo{
		Name: "jump_to_Home",
		Desc: "Jump to Home page",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"target_page": {
				Desc:     "The page to jump to",
				Type:     schema.String,
				Required: true,
			},
		}),
	}

	return utils.NewTool(info, jumpToHome)
}

type PageJumpParams struct {
	TargetPage string `json:"target_page"` // 用户请求跳转的页面
}

// 定义跳转处理函数
func jumpToHome(_ context.Context, params *PageJumpParams) (string, error) {
	return "/", nil
}

func createCartJumpTool() tool.InvokableTool {
	info := &schema.ToolInfo{
		Name: "jump_to_Cart",
		Desc: "Jump to Cart page",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"target_page": {
				Desc:     "The page to jump to",
				Type:     schema.String,
				Required: true,
			},
		}),
	}

	return utils.NewTool(info, jumpToCart)
}

func jumpToCart(_ context.Context, params *PageJumpParams) (string, error) {
	return "/cart", nil
}

func createOrderJumpTool() tool.InvokableTool {
	info := &schema.ToolInfo{
		Name: "jump_to_Order",
		Desc: "Jump to Order page",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"target_page": {
				Desc:     "The page to jump to",
				Type:     schema.String,
				Required: true,
			},
		}),
	}

	return utils.NewTool(info, jumpToOrder)
}

func jumpToOrder(_ context.Context, params *PageJumpParams) (string, error) {
	return "/order", nil
}
