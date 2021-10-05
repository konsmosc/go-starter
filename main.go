package main

import (
	"github.com/konsmosc/go-starter/database"
	config "github.com/konsmosc/go-starter/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	database.InitDb()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "up",
		})
	})

	r.Run(config.Port)
}
