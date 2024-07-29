package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"

	gen "oapi-codegen-with-middleware-example/generated"
)

// Global middleware
func globalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Global middleware: before request")
		c.Next()
		log.Println("Global middleware: after request")
	}
}

// Middleware that runs before the request
func beforeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Before request middleware")
		c.Next() // Pass control to the next handler
	}
}

// Middleware that runs after the request
func afterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Pass control to the next handler
		log.Println("After request middleware")
	}
}

// Middleware specific to  POST /items routes
func adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Admin-specific middleware")
		c.Next() // Pass control to the next handler
	}
}

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

func main() {
	router := gin.Default()

	server := &Server{}

	// Apply global middleware
	router.Use(globalMiddleware())

	// Create a route group with before and after middleware
	apiGroup := router.Group("/")
	{
		apiGroup.Use(beforeMiddleware())
		apiGroup.Use(afterMiddleware())
		// Register the generated server routes within the group
		gen.RegisterHandlersWithOptions(apiGroup, server, gen.GinServerOptions{})
	}

	adminGroup := router.Group("/items")
	{
		adminGroup.Use(adminMiddleware())
		gen.RegisterHandlersWithOptions(adminGroup, server, gen.GinServerOptions{})
	}

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
