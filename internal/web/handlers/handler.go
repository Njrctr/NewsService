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

	return router
}
