package main

import (
	"marcellinuselbert-be/controllers"
	"marcellinuselbert-be/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsFetch)
	r.GET("/posts/:id", controllers.PostsFetchById)
	r.Run() // listen and serve on 0.0.0.0:8080
}
