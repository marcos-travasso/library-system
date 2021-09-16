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

func setupUserFunctions() {
	router.POST("/users", postUser)
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.DELETE("/users/:id", deleteUser)
	router.PATCH("/users", updateUser)
}

func postUser(c *gin.Context) {
	var user structs.User

	if c.BindJSON(&user) == nil {
		id, err := dbDir.InsertUser(user)
		if err != nil {
			gotJSON, _ := json.Marshal(user)
			log.Printf("%s", err)
			log.Printf("%s", gotJSON)

			c.String(http.StatusBadRequest, "Failed to insert user")
			return
		}

		idString := fmt.Sprint(id)

		c.String(http.StatusOK, idString)
		return
	}

	c.String(http.StatusBadRequest, "Fail to parse JSON")
}

func getUsers(c *gin.Context) {
	users, err := dbDir.SelectUsers()
	if err != nil {
		log.Printf("%s", err)
		c.String(http.StatusInternalServerError, "Failed to select users")
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("%s", err)
		c.String(http.StatusBadRequest, "Fail to receive ID")
		return
	}
	receivedUser := structs.User{ID: id}

	err = dbDir.DeleteUser(receivedUser)
	if err != nil {
		gotJSON, _ := json.Marshal(receivedUser)
		log.Printf("%s", err)
		log.Printf("%s", gotJSON)

		c.String(http.StatusBadRequest, "Failed to delete user")
		return
	}

	c.String(http.StatusOK, "Success")
}

func getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("%s", err)
		c.String(http.StatusBadRequest, "Fail to receive ID")
		return
	}
	receivedUser := structs.User{ID: id}

	user, err := dbDir.SelectUser(receivedUser)
	if err != nil {
		log.Printf("%s", err)
		c.String(http.StatusInternalServerError, "Failed to select user")
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
	var user structs.User

	if c.BindJSON(&user) == nil {
		err := dbDir.UpdateUser(user)
		if err != nil {
			gotJSON, _ := json.Marshal(user)
			log.Printf("%s", err)
			log.Printf("%s", gotJSON)

			c.String(http.StatusBadRequest, "Failed to update user")
			return
		}

		c.String(http.StatusOK, "Success")
		return
	}

	c.String(http.StatusBadRequest, "Fail to parse JSON")
}
