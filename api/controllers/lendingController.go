package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/services"
	"log"
	"net/http"
	"strconv"
)

func initializeLendingController() {
	router.POST("/lendings", postLending)
	router.PATCH("/lendings/:id", returnLending)
}

func postLending(c *gin.Context) {
	var lending models.Lending

	if c.BindJSON(&lending) == nil {
		err := services.InsertLending(&lending)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, lending)
		return
	}

	log.Printf("postLending(): failed to parse JSON")
	c.String(http.StatusBadRequest, "failed to parse JSON")
}

func returnLending(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	receivedLending := models.Lending{ID: int64(id)}

	err = services.ReturnLending(&receivedLending)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "Success")
	return
}
