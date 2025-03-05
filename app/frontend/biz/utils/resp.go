package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v5"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	userId := frontendUtils.GetUserIdfromCtx(ctx)
	userRole := frontendUtils.GetUserRolefromCtx(ctx)
	content["user_id"] = userId
	content["user_role"] = userRole
	if userId > 0 {
		cartResp, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{
			UserId: uint32(userId),
		})
		if err == nil && cartResp != nil {
			content["cart_num"] = len(cartResp.Items)
		}
	}
	return content
}

func SaveFile(file io.Reader, destPath string) error {
	// 创建一个目标文件
	outFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer outFile.Close()

	// 使用 io.Copy 将内容从 file 写入 outFile
	_, err = io.Copy(outFile, file)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

// 生成 JWT Token
func GenerateJWT(userId int, role string) (string, error) {
	// JWT 密钥
	var mySigningKey = []byte(os.Getenv("JWT_SECRET"))

	// 创建 token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置 claims（负载信息）
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // 设置过期时间为 72 小时
	// 签署 JWT Token
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Base64UrlDecode 解码 Base64Url 编码的字符串
func base64UrlDecode(base64Url string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(base64Url)
}

type Payload struct {
	Exp    int    `json:"exp"`
	Role   string `json:"role"`
	UserID int    `json:"user_id"`
}

// ParseJWT 解析 JWT 并验证签名
func ParseJWT(jwtToken string) (header string, data Payload, err error) {
	// 分割 JWT
	parts := strings.Split(jwtToken, ".")
	if len(parts) != 3 {
		return "", data, fmt.Errorf("invalid JWT")
	}

	// 解码 Header 和 Payload
	headerBytes, err := base64UrlDecode(parts[0])
	if err != nil {
		return "", data, fmt.Errorf("failed to decode header: %v", err)
	}
	header = string(headerBytes)
	payloadBytes, err := base64UrlDecode(parts[1])
	if err != nil {
		return "", data, fmt.Errorf("failed to decode payload: %v", err)
	}
	// 输出 Header 和 Payload
	payload := string(payloadBytes)
	err = json.Unmarshal([]byte(payload), &data)
	if err != nil {
		return "", data, fmt.Errorf("failed to decode payload: %v", err)
	}
	// 验证签名
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", data, fmt.Errorf("failed to parse or validate JWT: %v", err)
	}

	// 如果 token 无效或签名不匹配
	if !token.Valid {
		return "", data, fmt.Errorf("invalid token signature")
	}
	return header, data, nil
}
