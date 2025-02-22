package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/healthcheck", func(c *gin.Context) {
		d, err := http.Get("http://localhost:2333/version")
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Health check failed",
			})
			return
		}
		if d.StatusCode != 200 {
			c.JSON(500, gin.H{
				"message": "Health check failed",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.Run()
}
