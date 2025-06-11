package handlers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {

	router := gin.Default()

	news := router.Group("/news")
	{
		news.GET("/", getNews)           // Получить новости по фильтру
		news.GET("/:id", getOneNews)     // Получить конкретную новость по id
		news.GET("/count", getNewsCount) // Получить количество новостей по фильтру
	}

	tags := router.Group("/tags")
	{
		tags.GET("/", getTags) // Получить все теги
	}

	categories := router.Group("/categories")
	{
		categories.GET("/", getCategories) // Получить все категории
	}

	return router
}
