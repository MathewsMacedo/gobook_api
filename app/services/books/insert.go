package books

import (
	"github.com/mathewsmacedo/go_api/app/models"
	"github.com/mathewsmacedo/go_api/db"
)

func Insert(input *models.CreateBookInput) models.Book {
	db := db.Connection()

	book := models.Book{Title: input.Title, Author: input.Author}

	db.Create(&book)

	return book
}
