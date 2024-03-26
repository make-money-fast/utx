package ginx

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// Param 获取去除了 / 的参数
func Param(ctx *gin.Context, name string) string {
	return strings.TrimPrefix(ctx.Param(name), "/")
}

// HtmlParam 获取伪静态的参数
func HtmlParam(ctx *gin.Context, name string) string {
	prefix := strings.TrimPrefix(ctx.Param(name), "/")
	return strings.TrimSuffix(prefix, ".html")
}
