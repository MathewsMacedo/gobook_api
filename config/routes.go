package config

import (
	"github.com/gin-gonic/gin"
	version1 "github.com/mathewsmacedo/go_api/app/controllers/api/v1"
	"github.com/mathewsmacedo/go_api/app/middleware"
)

func Routes() {
	r := gin.New()
	r.Use(middleware.CheckToken())

	api_v1 := r.Group("/api/v1")
	{
		books := api_v1.Group("/books")
		{
			books.GET("", version1.BookIndex)
			books.POST("", version1.BookCreate)
			books.GET("/:id", version1.BookShow)
			books.PATCH("/:id", version1.BookUpdate)
			books.DELETE("/:id", version1.BookDestroy)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Page not found"})
	})

	r.Run(":3000")
}
