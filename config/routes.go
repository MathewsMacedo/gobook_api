package config

import (
	"github.com/gin-gonic/gin"
	version1 "github.com/mathewsmacedo/go_api/app/controllers/api/v1"
)

func Routes() {
	r := gin.Default()

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

	r.Run(":3000")
}
