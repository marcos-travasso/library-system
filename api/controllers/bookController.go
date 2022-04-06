package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/services"
	"log"
	"net/http"
	"strconv"
)

func initializeBookController() {
	router.POST("/books", postBook)
	router.GET("/books/:id", getBook)
	//router.GET("/books", getBooks)
	//router.DELETE("/books/:id", deleteBook)
}

func postBook(c *gin.Context) {
	var book models.Book

	if c.BindJSON(&book) == nil {
		err := services.InsertBook(&book)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, book)
		return
	}

	log.Printf("failed to parse JSON")
	c.String(http.StatusBadRequest, "failed to parse JSON")
}

func getBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	receivedBook := models.Book{ID: int64(id)}

	err = services.SelectBook(&receivedBook)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, receivedBook)
}

//
//func getBooks(c *gin.Context) {
//	books, err := dbDir.SelectBooks()
//	if err != nil {
//		log.Printf("getBooks(): %s", err)
//
//		c.String(http.StatusInternalServerError, err.Error())
//		return
//	}
//	c.IndentedJSON(http.StatusOK, books)
//}
//
//func deleteBook(c *gin.Context) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		log.Printf("deleteBook(): %s", err)
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//	receivedBook := models.Book{ID: id}
//
//	err = dbDir.DeleteBook(receivedBook)
//	if err != nil {
//		gotJSON, _ := json.Marshal(receivedBook)
//		log.Printf("deleteBook(): %s", err)
//		log.Printf("deleteBook(): %s", gotJSON)
//
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	c.String(http.StatusOK, "Success")
//}
//
