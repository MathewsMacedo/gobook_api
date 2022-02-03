package books

import (
	"strconv"

	"github.com/mathewsmacedo/go_api/app/models"
	"github.com/mathewsmacedo/go_api/db"
)

func Update(input *models.UpdateBookInput) (models.Book, error) {
	book, err := Details(strconv.FormatInt(input.ID, 10))

	if err != nil {
		return book, err
	}

	db := db.Connection()

	db.Model(&book).Updates(input)

	return book, err
}
