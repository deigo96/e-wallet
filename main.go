package main

import (
	"net/http"

	"github.com/deigo96/e-wallet.git/app/handlers"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configuration := config.NewConfiguration()
	db := config.DBConnection(configuration.DbConfig)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders: []string{"X-Total-Count"},
	}))

	r.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group(configuration.APIVersion)
	handlers.NewHandler(configuration, db, v1)

	r.Run(configuration.ServiceHost + ":" + configuration.ServicePort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	defer config.CloseConnection(db)
}
