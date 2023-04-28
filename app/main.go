package main

import (
	"mnewsapi/app/configs"
	"mnewsapi/app/controllers"
	"mnewsapi/app/initializers"
	"mnewsapi/app/repository"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	db := configs.ConnectDB()

	newsRepository := repository.NewRepository(db)
	newsService := repository.NewService(newsRepository)
	newsHandler := controllers.NewNewsHandler(newsService)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", newsHandler.RootHandler)
	v1.GET("/news", newsHandler.GetNews)
	v1.GET("/news/:id", newsHandler.GetNew)
	v1.POST("/news", newsHandler.CreateNewsHandler)
	v1.PUT("/news/:id", newsHandler.UpdateNewsHandler)
	v1.DELETE("/news/:id", newsHandler.DeleteNewsHandler)

	router.Run()
}
