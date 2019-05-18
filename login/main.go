package main

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-web"
	"log"
	"net/http"
)

type User struct{}

func (u *User) Login(c *gin.Context) {
	log.Print("Received Login API request")
	var req map[string]interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&req)

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	if req["username"] == nil {
		err := errors.New("Field username is required.")
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	if req["password"] == nil {
		err := errors.New("Field password is required.")
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": true, "data": req})
}

func main() {
	service := web.NewService(
		web.Name("go.micro.api.login"),
	)

	service.Init()
	user := new(User)
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
		return
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
		return
	})

	router.POST("/login", user.Login)
	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
