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

	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	router.GET("/books", bookHandler.RootHandler)

	router.GET("/books/:id", bookHandler.GetBook)
	router.GET("/books/:id/:title", bookHandler.BooksHandler)
	// localhost:8080/query?id=1&title=bumi -> query with multiple params separated by &
	router.GET("/query", bookHandler.QueryHandler)
	router.POST("/books", bookHandler.PostBooksHandler)
	router.PUT("/books/:id", bookHandler.UpdateBook)
	router.DELETE("/books/:id", bookHandler.DeleteBook)

	//versioning router
	v1 := router.Group("/v1")
	v1.GET("/", bookHandler.RootHandler)

	v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	// localhost:8080/query?id=1&title=bumi -> query with multiple params separated by &
	v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)

	router.Run(":8080")

	//main
	//handler
	//service
	//repository
	//db
	//mysql
}
