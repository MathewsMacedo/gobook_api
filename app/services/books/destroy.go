package books

import (
	"github.com/mathewsmacedo/go_api/app/models"
	"github.com/mathewsmacedo/go_api/db"
)

func Destroy(id string) (models.Book, error) {
	db := db.Connection()

	var book models.Book

	err := db.Where("id = ?", id).First(&book).Error

	db.Delete(&book)

	return book, err
}
