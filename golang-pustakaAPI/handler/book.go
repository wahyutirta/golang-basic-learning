package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
	//bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var bookResponse []book.BookResponse

	for _, b := range books {
		res := book.BookResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Author:      b.Author,
			Price:       b.Price,
			Rating:      b.Rating,
			Discount:    b.Discount,
		}
		bookResponse = append(bookResponse, res)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

// value by params
func (h *bookHandler) BooksHandler(c *gin.Context) {
	title := c.Param("title")
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

// value by query
func (h *bookHandler) QueryHandler(c *gin.Context) {
	title := c.Query("title")
	id := c.Query("id")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}
