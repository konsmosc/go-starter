package main

import (
	"github.com/konsmosc/go-starter/database"
	// "github.com/konsmosc/go-starter/seed"
	config "github.com/konsmosc/go-starter/utils"

	"github.com/gin-contrib/cors"
	"github.com/konsmosc/go-starter/models"
	Routers "github.com/konsmosc/go-starter/routers"
)

func main() {
	r := Routers.Router()
	r.Use(cors.Default())

	database.InitDb()
	database.Db.AutoMigrate(&models.User{},
		&models.Score{},
		&models.Quiz{},
		&models.Question{},
		&models.Option{})

	// seed.Load(database.Db)

	r.Run(config.Port)
}
