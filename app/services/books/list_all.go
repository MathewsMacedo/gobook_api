package books

import (
	"github.com/mathewsmacedo/go_api/app/models"
	"github.com/mathewsmacedo/go_api/db"
)

func ListAll() []models.Book {
	var books []models.Book

	db := db.Connection()

	db.Find(&books)

	return books
}
