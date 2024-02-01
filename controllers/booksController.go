package controllers

import (
	"SampleCRM/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func getBooks(c *gin.Context) {
	var bookQuery models.Book
	if err := c.ShouldBindQuery(&bookQuery); err != nil {
		handleError(c, err, http.StatusBadRequest)
		return
	}
	books, err := models.SearchBooks(bookQuery)
	if err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, books)
}

func getBooks1(c *gin.Context) {
	var bookQuery models.Book
	if err := c.ShouldBindQuery(&bookQuery); err != nil {
		handleError(c, err, http.StatusBadRequest)
		return
	}
	books, err := models.SearchBooks(bookQuery)
	if err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, books)
}

func getBooks2(c *gin.Context) {
	var bookQuery models.Book
	if err := c.ShouldBindQuery(&bookQuery); err != nil {
		handleError(c, err, http.StatusBadRequest)
		return
	}
	books, err := models.SearchBooks(bookQuery)
	if err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		handleError(c, err, http.StatusBadRequest)
		return
	}
	if err := models.CreateBook(&book); err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, book)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		handleError(c, err, http.StatusBadRequest)
		return
	}
	if err := models.UpdateBook(id, &book); err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book Updated successfully"})

}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := models.DeleteBook(id); err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book Deleted successfully"})
}

func handleError(c *gin.Context, err error, statusCode int) {
	log.Println(err)
	c.AbortWithStatusJSON(statusCode, err)
}

func RegisterBookRoutes(r *gin.Engine) {
	subRouter := r.Group("/books")
	{
		subRouter.GET("/", getBooks)
		subRouter.POST("/", createBook)
		subRouter.PUT("/:id", updateBook)
		subRouter.DELETE("/:id", deleteBook)
	}
}
