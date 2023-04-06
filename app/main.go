package main

import (
	"log"
	"mnewsapi/app/handler"
	"mnewsapi/app/initializers"
	"mnewsapi/app/news"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error")
	}

	/// Auto Migrate (Add field table using table berdasarkan nama type Struct)
	db.AutoMigrate(&news.News{})

	newsRepository := news.NewRepository(db)
	newsService := news.NewService(newsRepository)
	newsHandler := handler.NewNewsHandler(newsService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", newsHandler.RootHandler)
	v1.GET("/hello", newsHandler.HelloHandler)

	v1.GET("/news/:title/:desc", newsHandler.NewsHandler) //Path Variable
	v1.GET("/query", newsHandler.QueryHandler)            //Query String (Params)

	v1.POST("/news", newsHandler.PostNewsHandler)

	v2 := router.Group("/v2")
	v2.GET("/")
	v2.GET("/news", newsHandler.GetNews)
	v2.GET("/news/:id", newsHandler.GetNew)
	v2.POST("/news", newsHandler.CreateNewsHandler)
	v2.PUT("/news/:id", newsHandler.UpdateNewsHandler)
	v2.DELETE("/news/:id", newsHandler.DeleteNewsHandler)

	router.Run()
}
