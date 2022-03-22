package main

import (
	"blog/infrastructure"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		fmt.Println("Hoola holla holla")
		log.Println("Hoola holla hollads")

		infrastructure.LoadEnv()     //loading env
		infrastructure.NewDatabase() //new database connection

		context.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})

	router.Run(":8000")
}
