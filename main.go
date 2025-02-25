package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Token string `yaml:"token"`
	Port  int64  `yaml:"port"`
	Url   string `yaml:"url"`
}

func main() {
	r := gin.Default()
	client := &http.Client{
		Timeout: 5 * 60,
	}

	yamlfile, err := os.ReadFile("set.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlfile, &config)
	if err != nil {
		panic(err)
	}
	r.GET("/healthcheck", func(c *gin.Context) {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s:%d/version", config.Url, config.Port), nil)
		if err != nil {
			c.JSON(500, gin.H{"message": "Failed to create request"})
			return
		}
		req.Header.Set("Authorization", config.Token)

		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != 200 {
			c.JSON(500, gin.H{"message": "Version check failed"})
			return
		}
		defer resp.Body.Close()
		req, err = http.NewRequest("GET", fmt.Sprintf("%s:%d/v4/loadtracks?identifier=dQw4w9WgXcQ", config.Url, config.Port), nil)
		if err != nil {
			c.JSON(500, gin.H{"message": "Failed to create request"})
			return
		}
		req.Header.Set("Authorization", config.Token)

		resp, err = client.Do(req)
		if err != nil || resp.StatusCode != 200 {
			c.JSON(500, gin.H{"message": "Load tracks check failed"})
			return
		}
		defer resp.Body.Close()

		c.JSON(200, gin.H{"message": "ok"})
	})
	r.Run()
}
