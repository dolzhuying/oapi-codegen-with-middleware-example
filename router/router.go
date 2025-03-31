package router

import (
    "oapi-codegen-with-middleware-example/generated"
    "oapi-codegen-with-middleware-example/handlers"
    "oapi-codegen-with-middleware-example/middlewares"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    
    // 全局中间件
    router.Use(middlewares.GlobalMiddleware)
    
    // API分组使用特定中间件
    apiGroup := router.Group("/api")
    apiGroup.Use(
        middlewares.AuthMiddleware,
        middlewares.LoggingMiddleware,
    )
    
    //为实现了ServerInterface的server注册路由
    server:=&handlers.Server{}
    openapi.RegisterHandlers(apiGroup, server)
    
    return router
}