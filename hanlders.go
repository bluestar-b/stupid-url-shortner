package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func addURL(c *gin.Context) {
	var payload struct {
		LongURL string `json:"longURL"`
	}

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	longURL := payload.LongURL
	urlID := generateURLID(longURL)
	createdAt := time.Now()

	mutex.Lock()
	db[urlID] = URLInfo{LongURL: longURL, UserAgent: c.GetHeader("User-Agent"), CreatedAt: createdAt}
	mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{"urlID": urlID, "createdAt": createdAt})

	saveDataToJSONFile()
}
