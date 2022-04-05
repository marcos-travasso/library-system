package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-travasso/library-system/models"
	"github.com/marcos-travasso/library-system/services"
	"log"
	"net/http"
)

func initializeUserController() {
	router.POST("/users", postUser)
	//router.GET("/users", getUsers)
	//router.GET("/users/:id", getUser)
	//router.DELETE("/users/:id", deleteUser)
	//router.PATCH("/users", updateUser)
}

func postUser(c *gin.Context) {
	var user models.User

	if c.BindJSON(&user) == nil {
		err := services.InsertUser(&user)
		//TODO better errors check issue #9
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, user)
		return
	}

	log.Printf("failed to parse JSON")
	c.String(http.StatusBadRequest, "failed to parse JSON")
}

//
//func getUsers(c *gin.Context) {
//	users, err := dbDir.SelectUsers()
//	if err != nil {
//		log.Printf("getUsers(): %s", err)
//
//		c.String(http.StatusInternalServerError, err.Error())
//		return
//	}
//	c.IndentedJSON(http.StatusOK, users)
//}
//
//func deleteUser(c *gin.Context) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		log.Printf("deleteUser(): %s", err)
//
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//	receivedUser := models.User{ID: id}
//
//	err = dbDir.DeleteUser(receivedUser)
//	if err != nil {
//		gotJSON, _ := json.Marshal(receivedUser)
//		log.Printf("deleteUser(): %s", err)
//		log.Printf("deleteUser(): %s", gotJSON)
//
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	c.String(http.StatusOK, "Success")
//}
//
//func getUser(c *gin.Context) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		log.Printf("getUser(): %s", err)
//
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//	receivedUser := models.User{ID: id}
//
//	user, err := dbDir.SelectUser(receivedUser)
//	if err != nil {
//		log.Printf("getUser(): %s", err)
//
//		c.String(http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	c.IndentedJSON(http.StatusOK, user)
//}
//
//func updateUser(c *gin.Context) {
//	var user models.User
//
//	if c.BindJSON(&user) == nil {
//		err := dbDir.UpdateUser(user)
//		if err != nil {
//			gotJSON, _ := json.Marshal(user)
//			log.Printf("updateUser(): %s", err)
//			log.Printf("updateUser(): %s", gotJSON)
//
//			c.String(http.StatusBadRequest, err.Error())
//			return
//		}
//
//		c.String(http.StatusOK, "Success")
//		return
//	}
//
//	log.Printf("updateUser(): failed to parse JSON")
//	c.String(http.StatusBadRequest, "Fail to parse JSON")
//}
