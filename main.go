package main

import (
	"github.com/OctavianoRyan25/VhiWEB/config"
	"github.com/OctavianoRyan25/VhiWEB/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	config.AutoMigrate()

	r := gin.Default()
	routes.InitRoutes(r)
	r.Run(":8080")
}
