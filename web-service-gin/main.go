package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{
		ID:     "1",
		Name:   "Harry Potter",
		Author: "J.K. Rowling",
		Price:  15.9,
	},
	{
		ID:     "2",
		Name:   "One Piece",
		Author: "Oda Eiichirō",
		Price:  2.99,
	},
	{
		ID:     "3",
		Name:   "demon slayer",
		Author: "koyoharu gotouge",
		Price:  2.99,
	},
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/book/:id", getBookByID)
	router.POST("/book", createbook)
	router.PUT("/book/:id", UpdateBookByID)

	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {

	paramID := c.Param("id")
	for _, book := range books {
		if paramID == book.ID {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, "book not found")

}

func createbook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func UpdateBookByID(c *gin.Context) {
	var updateBook book

	if err := c.BindJSON(&updateBook); err != nil {
		return
	}

	paramID := c.Param("id")

	for i := range books {
		if paramID == books[i].ID {

			books[i].Name = updateBook.Name
			books[i].Author = updateBook.Author
			books[i].Price = updateBook.Price

			c.JSON(http.StatusOK, books[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, "book not found")

}
