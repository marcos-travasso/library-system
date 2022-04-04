package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marcos-travasso/library-system/models"
	"log"
	"net/http"
	"strconv"
)

func setupBookEndpoints() {
	router.POST("/books", postBook)
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.DELETE("/books/:id", deleteBook)
	router.PATCH("/books", updateBook)
}

func postBook(c *gin.Context) {
	var book models.Book

	if c.BindJSON(&book) == nil {
		id, err := dbDir.InsertBook(book)
		if err != nil {
			gotJSON, _ := json.Marshal(book)
			log.Printf("postBook(): %s", err)
			log.Printf("postBook(): %s", gotJSON)

			c.String(http.StatusBadRequest, err.Error())
			return
		}

		idString := fmt.Sprint(id)

		c.String(http.StatusOK, idString)
		return
	}

	log.Printf("postBook(): failed to parse JSON")
	c.String(http.StatusBadRequest, "failed to parse JSON")
}

func getBooks(c *gin.Context) {
	books, err := dbDir.SelectBooks()
	if err != nil {
		log.Printf("getBooks(): %s", err)

		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("deleteBook(): %s", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	receivedBook := models.Book{ID: id}

	err = dbDir.DeleteBook(receivedBook)
	if err != nil {
		gotJSON, _ := json.Marshal(receivedBook)
		log.Printf("deleteBook(): %s", err)
		log.Printf("deleteBook(): %s", gotJSON)

		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, "Success")
}

func getBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("getBook(): %s", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	receivedBook := models.Book{ID: id}

	book, err := dbDir.SelectBook(receivedBook)
	if err != nil {
		gotJSON, _ := json.Marshal(book)
		log.Printf("getBook(): %s", err)
		log.Printf("getBook(): %s", gotJSON)

		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func updateBook(c *gin.Context) {
	var book models.Book

	if c.BindJSON(&book) == nil {
		err := dbDir.UpdateBook(book)
		if err != nil {
			gotJSON, _ := json.Marshal(book)
			log.Printf("updateBook(): %s", err)
			log.Printf("updateBook(): %s", gotJSON)

			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.String(http.StatusOK, "Success")
		return
	}

	log.Printf("postBook(): failed to parse JSON")
	c.String(http.StatusBadRequest, "failed to parse JSON")
}
