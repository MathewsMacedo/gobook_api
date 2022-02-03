package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mathewsmacedo/go_api/app/models"
	"github.com/mathewsmacedo/go_api/app/services/books"
)

func BookIndex(c *gin.Context) {
	books := books.ListAll()
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func BookCreate(c *gin.Context) {
	var input models.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := books.Insert(&input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func BookShow(c *gin.Context) {
	book, err := books.Details(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func BookUpdate(c *gin.Context) {
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

func BookDestroy(c *gin.Context) {
	_, err := books.Destroy(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
