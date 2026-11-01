package main

import (
	"ciclobits/internal/bootstrap"
	"ciclobits/internal/db"
	"ciclobits/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loanding .env file")
	}	
}


func main() {

	db := db.NewDb()
	app := bootstrap.BuiltApp(db)
	r := gin.Default()
	routes.InitRoutes(app, r)

	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080
}