package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "admin",
		"bio":  "A sistem Admins",
	})
}
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "Belajar Golang",
		"bio":   "Belajar API Golang euui",
	})
}

// value by params
func BooksHandler(c *gin.Context) {
	title := c.Param("title")
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

// value by query
func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	id := c.Query("id")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func PostBooksHandler(c *gin.Context) {
	var book book.BookInput

	err := c.ShouldBindJSON(&book)

	if err != nil {
		// sending error message as json
		erroMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			erroMessage := fmt.Sprintf("Error on fields %s, condition: %s", e.Field(), e.ActualTag())
			erroMessages = append(erroMessages, erroMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": erroMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title":  book.Title,
		"price":  book.Price,
		"author": book.Author,
	})

}
