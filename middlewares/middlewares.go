package middlewares

import (
	"log"
	"net/http"


	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware()gin.HandlerFunc {
    return func(ctx *gin.Context){
		token:=ctx.GetHeader("Authorization")
		if token==""{
			ctx.JSON(http.StatusUnauthorized,gin.H{"error":"missing authorization"})
			ctx.Abort()
			return
		}

		// username,err:=utils.ParseJWT(token)
		// if err!=nil{
		// 	ctx.JSON(http.StatusUnauthorized,gin.H{"error":err.Error()})
		// 	ctx.Abort()
		// 	return
		// }

		ctx.Set("username","hello")
		ctx.Next()//下一步调用
	}
}

func LoggingMiddleware()gin.HandlerFunc {
    return func(ctx *gin.Context){
        log.Printf("请求: %s %s | 状态: %d | 耗时: %v",
        ctx.Request.Method, ctx.Request.URL.Path, 200, 100)
        ctx.Next()
	}
   
}


func GlobalMiddleware()gin.HandlerFunc {
    return func(ctx *gin.Context){
        ctx.Next()
    }
}