package main

import (
	"github.com/LordOfNightmares/Go-Testing/api"
	"github.com/LordOfNightmares/Go-Testing/lib/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	//load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	//
	//// initializes database
	//db, _ := database.Initialize()
	//

	port := os.Getenv("PORT")
	app := gin.Default() // create gin app
	//app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())
	api.ApplyRoutes(app)      // apply api router
	err = app.Run(":" + port) // listen to given port
}
