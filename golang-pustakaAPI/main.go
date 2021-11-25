package main

import (
	"fmt"
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
	router := gin.Default()

	//  CRUD

	// Create data
	// book := book.Book{}

	// book.Title = "Mewarnai Rusa"
	// book.Description = "Ini Adalah Buku Mewarnai Rusa Jantan dan Betina"
	// book.Price = 500000
	// book.Rating = 5
	// book.Discount = 10

	// err = db.Create(&book).Error

	// if err != nil {
	// 	fmt.Println("======================================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("======================================")
	// }

	// one book
	// var book book.Book
	// err = db.Debug().First(&book, 1).Error

	// if err != nil {
	// 	fmt.Println("======================================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("======================================")
	// }
	// fmt.Printf("book title: %v", book.Title)
	// fmt.Printf("book object: %v", book)

	// all book
	// var books []book.Book
	// err = db.Debug().Find(&books).Error

	// if err != nil {
	// 	fmt.Println("======================================")
	// 	fmt.Println("Error finding all book record")
	// 	fmt.Println("======================================")
	// }

	// for _, b := range books {
	// 	fmt.Printf("book title: %v", b.Title)
	// 	fmt.Printf("book object: %v", b)
	// }

	// find book with condition
	// var book book.Book

	// err = db.Debug().Where("Title = ?", "Mewarnai Rusa").First(&book).Error  // -> find one
	// err = db.Debug().Where("Title = ?", "Mewarnai Rusa").Find(&book).Error

	// if err != nil {
	// 	fmt.Println("======================================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("======================================")
	// }
	// fmt.Printf("book title: %v\n", book.Title)
	// fmt.Printf("book object: %v", book)

	// update book
	var book book.Book
	err = db.Debug().Where("Id = ?", "1").Find(&book).Error

	if err != nil {
		fmt.Println("======================================")
		fmt.Println("Error finding book record")
		fmt.Println("======================================")
	}
	book.Title = "Mewarnai Rusa (Revised Edition)"

	err = db.Save(&book).Error

	if err != nil {
		fmt.Println("======================================")
		fmt.Println("Error finding book record")
		fmt.Println("======================================")
	}
	fmt.Printf("book title: %v\n", book.Title)
	fmt.Printf("book object: %v", book)

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
