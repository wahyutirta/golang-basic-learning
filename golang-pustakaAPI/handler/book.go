package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

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

func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	b, err := h.bookService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookRespose(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

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
		res := convertToBookRespose(b)
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

func (h *bookHandler) UpdateBook(c *gin.Context) {
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

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)

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

func convertToBookRespose(bookObj book.Book) book.BookResponse {

	return book.BookResponse{
		ID:          bookObj.ID,
		Title:       bookObj.Title,
		Description: bookObj.Description,
		Author:      bookObj.Author,
		Price:       bookObj.Price,
		Rating:      bookObj.Rating,
		Discount:    bookObj.Discount,
	}
}
