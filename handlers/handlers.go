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
	items := []gen.Item{
		{Id: null.StringFrom("1").Ptr(), Name: null.StringFrom("Item 1").Ptr(), Description: null.StringFrom("Description for Item 1").Ptr()},
		{Id: null.StringFrom("2").Ptr(), Name: null.StringFrom("Item 2").Ptr(), Description: null.StringFrom("Description for Item 2").Ptr()},
	}
	ctx.JSON(http.StatusOK, items)
}

// CreateItem is the implementation of the POST /items endpoint
func (s *Server) PostItems(ctx *gin.Context) {
	var newItem gen.NewItem
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