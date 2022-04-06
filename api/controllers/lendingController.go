package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/services"
	"log"
	"net/http"
)

func initializeLendingController() {
	router.POST("/lendings", postLending)
	//router.PATCH("/lendings/:id", returnLending)
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

//
//func getLendings(c *gin.Context) {
//	lendings, err := dbDir.SelectLendings()
//	if err != nil {
//		log.Printf("getLendings(): %s", err)
//
//		c.String(http.StatusInternalServerError, err.Error())
//		return
//	}
//	c.IndentedJSON(http.StatusOK, lendings)
//}
//
//
//	c.IndentedJSON(http.StatusOK, lending)
//}
//
//func returnLending(c *gin.Context) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		log.Printf("returnLending(): %s", err)
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//	receivedLending := models.Lending{ID: id}
//
//	err = dbDir.ReturnBook(receivedLending)
//	if err != nil {
//		gotJSON, _ := json.Marshal(receivedLending)
//		log.Printf("returnLending(): %s", err)
//		log.Printf("returnLending(): %s", gotJSON)
//
//		c.String(http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	c.IndentedJSON(http.StatusOK, "Success")
//}
