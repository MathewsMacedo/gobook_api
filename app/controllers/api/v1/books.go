package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mathewsmacedo/go_api/app/models"
	"github.com/mathewsmacedo/go_api/app/services/books"
)

// @BasePath /api/v1/books

// Books swagger
// @Summary Index
// @Schemes
// @Description List All Books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} models.Book
// @Router / [get]
func IndexBook(c *gin.Context) {
	books := books.ListAll()
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// @Summary Create
// @Schemes
// @Description Create Book
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} models.Book
// @Router / [post]
func CreateBook(c *gin.Context) {
	var input models.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := books.Insert(&input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// @Summary Show
// @Schemes
// @Description Show Book
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} models.Book
// @Router /:id [get]
func ShowBook(c *gin.Context) {
	book, err := books.Details(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// @Summary Update
// @Schemes
// @Description Update Book
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} models.Book
// @Router /:id [patch]
func UpdateBook(c *gin.Context) {
	// Validate input
	var input models.UpdateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	input.ID = id

	book, err := books.Update(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// @Summary Destroy
// @Schemes
// @Description Delete Book
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} models.Book
// @Router /:id [delete]
func DestroyBook(c *gin.Context) {
	_, err := books.Destroy(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
