package main

import (
	"github.com/gin-gonic/gin"
	"github.com/girishyt/go-crud/controllers"
	"github.com/girishyt/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostById)
	r.PUT("/posts/:id", controllers.UpdatePostById)
	r.DELETE("/posts/:id", controllers.DeletePostById)
	r.Run() // listen and serve on 0.0.0.0:8080
}
