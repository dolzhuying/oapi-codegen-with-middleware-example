package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token == "" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
            "error": "未授权访问",
        })
        return
    }
    // 验证token逻辑...
    
    // 添加信息到上下文中供后续处理器使用
    c.Set("userID", "用户ID值")
    
    // 继续下一个中间件或处理器
}

// LoggingMiddleware 日志中间件
func LoggingMiddleware(c *gin.Context) {
    // 请求开始前的逻辑
    startTime := time.Now()
    
    // 继续处理请求
    c.Next()
    
    // 请求结束后的逻辑
    duration := time.Since(startTime)
    statusCode := c.Writer.Status()
    
    // 记录日志
    log.Printf("请求: %s %s | 状态: %d | 耗时: %v",
        c.Request.Method, c.Request.URL.Path, statusCode, duration)
}

// ErrorHandlerMiddleware 错误处理中间件
func ErrorHandlerMiddleware(c *gin.Context, err error, statusCode int) {
    // 根据不同错误码返回标准化响应
    switch statusCode {
    case http.StatusNotFound:
        c.JSON(statusCode, gin.H{"error": "未找到资源", "details": err.Error()})
    case http.StatusBadRequest:
        c.JSON(statusCode, gin.H{"error": "请求参数错误", "details": err.Error()})
    case http.StatusInternalServerError:
        c.JSON(statusCode, gin.H{"error": "服务器内部错误", "details": err.Error()})
    default:
        c.JSON(statusCode, gin.H{"error": err.Error()})
    }
}

func GlobalMiddleware(c *gin.Context) {
    log.Printf("hello")
}