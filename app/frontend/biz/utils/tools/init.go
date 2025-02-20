/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tools

import (
	"context"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

var (
	Agent compose.Runnable[[]*schema.Message, []*schema.Message]
)

func Ai_init() {
	//openAIAPIKey := os.Getenv("OPENAI_API_KEY")

	ctx := context.Background()

	// 创建 Google Search 工具
	searchTool, err := duckduckgo.NewTool(ctx, &duckduckgo.Config{})
	if err != nil {
		Errorf("NewGoogleSearchTool failed, err=%v", err)
		return
	}

	// 初始化 tools
	todoTools := []tool.BaseTool{
		createHomeJumpTool(),
		createCartJumpTool(),
		createOrderJumpTool(),
		awakeClodeWeGo(),
		searchTool,
	}

	// 创建并配置 ChatModel
	chatModel, err := ollama.NewChatModel(context.Background(), &ollama.ChatModelConfig{
		BaseURL: "http://localhost:11434", // Ollama 服务地址
		Model:   "llama3.2",               // 模型名称
	})
	if err != nil {
		Errorf("NewChatModel failed, err=%v", err)
		return
	}

	// 获取工具信息, 用于绑定到 ChatModel
	toolInfos := make([]*schema.ToolInfo, 0, len(todoTools))
	var info *schema.ToolInfo
	for _, todoTool := range todoTools {
		info, err = todoTool.Info(ctx)
		if err != nil {
			Infof("get ToolInfo failed, err=%v", err)
			return
		}
		toolInfos = append(toolInfos, info)
	}

	// 将 tools 绑定到 ChatModel
	err = chatModel.BindTools(toolInfos)
	if err != nil {
		Errorf("BindTools failed, err=%v", err)
		return
	}

	// 创建 tools 节点
	todoToolsNode, err := compose.NewToolNode(context.Background(), &compose.ToolsNodeConfig{
		Tools: todoTools,
	})
	if err != nil {
		Errorf("NewToolNode failed, err=%v", err)
		return
	}

	// 构建完整的处理链
	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(chatModel, compose.WithNodeName("chat_model")).
		AppendToolsNode(todoToolsNode, compose.WithNodeName("tools"))

	// 编译并运行 chain
	Agent, err = chain.Compile(ctx)
	if err != nil {
		Errorf("chain.Compile failed, err=%v", err)
		return
	}

}
