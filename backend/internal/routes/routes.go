package routes

import (
	"ciclobits/internal/bootstrap"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *bootstrap.App, r *gin.Engine){
	
	api := r.Group("api").Group("v1")
	
	//Vamos definir as rotas na hora de criar o contrato
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}