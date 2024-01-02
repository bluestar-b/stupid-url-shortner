package main

import (
	"math/rand"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	loadDataFromJSONFile()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST"},
		AllowHeaders:  []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	router.POST("/api/add", addURL)
	router.GET("/:urlID", redirectToLongURL)

	router.Run(":8653")

	saveDataToJSONFile()
}
