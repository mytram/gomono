package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	app "github.com/mytram/gomono/tetra/internal/app/handlers"
)

func main() {
	os.Setenv("TZ", "UTC")

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
	}))

	app.Draw(r)

	r.Run()
	log.Println("tetra is running")
}
