package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// gorm db
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)

	book := book.BookRequest{
		Title:       "Mewarnai Kelinci",
		Author:      "Dendi Simanjuntak",
		Description: "Ini Adalah Buku Mewarnai Kelinci Jantan dan Betina untuk Anak Balita",
		Price:       "500000",
		Rating:      "5",
		Discount:    "10",
	}

	bookService.Create(book)

	router := gin.Default()

	router.GET("/", handler.RootHandler)
	router.GET("/hello", handler.HelloHandler)
	router.GET("/books/:id/:title", handler.BooksHandler)
	// localhost:8080/query?id=1&title=bumi -> query with multiple params separated by &
	router.GET("/query", handler.QueryHandler)
	router.POST("/books", handler.PostBooksHandler)

	//versioning router
	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	// localhost:8080/query?id=1&title=bumi -> query with multiple params separated by &
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8080")

}
