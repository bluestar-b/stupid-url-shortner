package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func redirectToLongURL(c *gin.Context) {
	urlID := c.Param("urlID")
	raw := c.Query("raw")

	mutex.Lock()
	urlInfo, exists := db[urlID]
	mutex.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	if raw == "true" {
		c.JSON(http.StatusOK, urlInfo)
		return
	}

	c.Redirect(http.StatusMovedPermanently, urlInfo.LongURL)
}
