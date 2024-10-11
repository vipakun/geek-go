package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, world")
	})

	// http://localhost:8080/users/xiangkun
	server.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "Hello, "+name)
	})

	// GET /order?id=123
	server.GET("/order", func(ctx *gin.Context) {
		id := ctx.Query("id")
		ctx.String(http.StatusOK, "ordering, "+id)
	})

	// http://localhost:8080/views/aasdg.html
	// return ordering, /aasdg.html
	server.GET("/views/*.html", func(ctx *gin.Context) {
		view := ctx.Param(".html")
		ctx.String(http.StatusOK, "ordering, "+view)
	})

	server.POST("/login", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, login")
	})
	server.Run(":8080")
}
