package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func MakeImageURL(image []byte, extension string) (string, error) {
	// 设置图片保存的目录
	saveDir := "./uploads"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		log.Printf("Error while creating directory: %v", err)
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// 使用当前时间戳生成唯一文件名，避免文件名冲突
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension) // 例如：1628865573502731000.jpg
	savePath := filepath.Join(saveDir, fileName)

	// 创建文件并将图片二进制数据写入文件
	outFile, err := os.Create(savePath)
	if err != nil {
		log.Printf("Error while creating file: %v", err)
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	// 将二进制数据写入文件
	_, err = outFile.Write(image)
	if err != nil {
		log.Printf("Error while writing file: %v", err)
		return "", fmt.Errorf("failed to write file: %v", err)
	}

	// 返回保存的文件路径
	return savePath, nil
}
