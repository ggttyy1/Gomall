package tools

import (
	"fmt"
	"time"
)

const (
	// Color codes for terminal output
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorBrown = "\033[31;1m"
	colorReset = "\033[0m"
)

func Errorf(format string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	prefix := fmt.Sprintf("%s[ERROR] %s ", colorRed, timestamp)
	message := fmt.Sprintf(format, args...)
	fmt.Printf("%s%s%s\n", prefix, message, colorReset)
}

func Infof(format string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	prefix := fmt.Sprintf("%s[INFO] %s ", colorGreen, timestamp)
	message := fmt.Sprintf(format, args...)
	fmt.Printf("%s%s%s\n", prefix, message, colorReset)
}
