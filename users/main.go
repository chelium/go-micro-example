package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-web"
	"log"
	"net/http"
)

type User struct{}

var users = json.RawMessage(`[{"username": "che", "email": "che@chelium.com"},{"username": "che2", "email": "che2@chelium.com"}]`)

func (u *User) List(c *gin.Context) {
	log.Print("Received User.List API request")
	c.JSON(http.StatusOK, gin.H{"result": true, "data": users})
}

func main() {
	service := web.NewService(
		web.Name("go.micro.api.user"),
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

	router.GET("/user/list", user.List)
	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
