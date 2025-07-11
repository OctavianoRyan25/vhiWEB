package routes

import (
	"github.com/OctavianoRyan25/VhiWEB/controller"
	"github.com/OctavianoRyan25/VhiWEB/controller/auth"
	"github.com/OctavianoRyan25/VhiWEB/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {

	api := r.Group("/api/v1")
	api.POST("/register", auth.Register)
	api.POST("/login", auth.Login)

	vendor := api.Group("/vendor")
	vendor.Use(middleware.AuthMiddleware())
	{
		vendor.POST("/create", controller.CreateVendor)
	}

	catalog := api.Group("/catalog")
	catalog.Use(middleware.AuthMiddleware())
	{
		catalog.POST("/create", controller.CreateCatalog)
		catalog.GET("/", controller.GetAllCatalogs)
		catalog.GET("/:id", controller.GetCatalogByID)
		catalog.PUT("/:id", controller.UpdateCatalog)
		catalog.DELETE("/:id", controller.DeleteCatalog)
	}
}
