package handlers

import (
    "net/http"
    "github.com/guregu/null"
    gen "oapi-codegen-with-middleware-example/generated"

    "github.com/gin-gonic/gin"
)

// Implementing the ServerInterface defined in the generated code
type Server struct{}

// GetItems is the implementation of the GET /items endpoint
func (s *Server) GetItems(ctx *gin.Context) {
	var item gen.NewItem
	//解析请求参数（通过gin框架ShouldBindJSON方法将请求参数绑定到item）
	//
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//封装响应数据
	resp:=gen.Response{
		Code:    200,
		Message: "Bad Request",
	}
	//write方法写入字节数据，直接对writer调用方法较为冗余
	ctx.Writer.WriteHeader(resp.Code)
	ctx.Writer.Write([]byte(resp.Message))

	//调用json方法自动设置正确的响应头，自动设置正确的响应头，自动设置正确的响应头，自动设置正确的响应头
	ctx.JSON(http.StatusOK, &resp)
	
}

// CreateItem is the implementation of the POST /items endpoint
func (s *Server) PostItems(ctx *gin.Context) {
	var newItem gen.NewItem
	//解析请求参数
	if err := ctx.ShouldBindJSON(&newItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item := gen.Item{
		Id:          null.StringFrom("3").Ptr(), // In a real app, you'd generate a unique ID
		Name:        newItem.Name,
		Description: newItem.Description,
	}
	ctx.JSON(http.StatusCreated, item)
}

// GetItemById is the implementation of the GET /items/{id} endpoint
func (s *Server) GetItemsId(ctx *gin.Context, id string) {
	if id == "1" {
		item := gen.Item{Id: null.StringFrom("1").Ptr(), Name: null.StringFrom("Item 1").Ptr(), Description: null.StringFrom("Description for Item 1").Ptr()}
		ctx.JSON(http.StatusOK, item)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	}
}