package config

import (
	"github.com/gin-gonic/gin"
	version1 "github.com/mathewsmacedo/go_api/app/controllers/api/v1"
	"github.com/mathewsmacedo/go_api/app/middleware"
	"github.com/mathewsmacedo/go_api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes() {
	r := gin.New()
	r.Use(middleware.CheckToken())

	docs.SwaggerInfo.BasePath = "/api/v1"

	api_v1 := r.Group("/api/v1")
	{
		books := api_v1.Group("/books")
		{
			books.GET("", version1.IndexBook)
			books.POST("", version1.CreateBook)
			books.GET("/:id", version1.ShowBook)
			books.PATCH("/:id", version1.UpdateBook)
			books.DELETE("/:id", version1.DestroyBook)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Page not found"})
	})

	r.Run(":3000")
}
