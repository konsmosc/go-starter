package Routers

import (
	"github.com/gin-gonic/gin"
	"github.com/konsmosc/go-starter/controllers"
)

func Router() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/status", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "up",
			})
		})
		api.GET("/user/:id", controllers.GetUser)
		api.PUT("/user/:id", controllers.UpdateUser)
		api.POST("/signup", controllers.SignUp)
		api.POST("/login", controllers.Login)
		api.GET("/quiz/:id", controllers.GetQuestions)
		api.POST("/submit", controllers.SubmitQuiz)
		api.GET("/performance/:id", controllers.GetPerformance)
	}

	return r
}
