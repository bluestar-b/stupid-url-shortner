package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	loadDataFromJSONFile()

	router := gin.Default()

	router.POST("/api/add", addURL)
	router.GET("/:urlID", redirectToLongURL)

	router.Run(":8080")

	saveDataToJSONFile()
}
