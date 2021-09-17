package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marcos-travasso/library-system/api/structs"
	"log"
	"net/http"
	"strconv"
)

func setupLendingEndpoints() {
	router.POST("/lendings", postLending)
	router.GET("/lendings", getLendings)
	router.GET("/lendings/:id", getLending)
	router.PATCH("/lendings/:id", returnLending)
}

func postLending(c *gin.Context) {
	var lending structs.Lending

	if c.BindJSON(&lending) == nil {
		id, err := dbDir.InsertLending(lending)
		if err != nil {
			gotJSON, _ := json.Marshal(lending)
			log.Printf("%s", err)
			log.Printf("%s", gotJSON)

			c.String(http.StatusBadRequest, err.Error())
			return
		}

		idString := fmt.Sprint(id)

		c.String(http.StatusOK, idString)
		return
	}

	c.String(http.StatusBadRequest, "failed to parse JSON")
}

func getLendings(c *gin.Context) {
	lendings, err := dbDir.SelectLendings()
	if err != nil {
		log.Printf("%s", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, lendings)
}

func getLending(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("%s", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	receivedLending := structs.Lending{ID: id}

	lending, err := dbDir.SelectLending(receivedLending)
	if err != nil {
		log.Printf("%s", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, lending)
}

func returnLending(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("%s", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	receivedLending := structs.Lending{ID: id}

	err = dbDir.ReturnBook(receivedLending)
	if err != nil {
		log.Printf("%s", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "Success")
}
