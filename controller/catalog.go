package controller

import (
	"github.com/OctavianoRyan25/VhiWEB/config"
	"github.com/OctavianoRyan25/VhiWEB/model"
	"github.com/OctavianoRyan25/VhiWEB/res"
	"github.com/OctavianoRyan25/VhiWEB/util"
	"github.com/gin-gonic/gin"
)

func CreateCatalog(c *gin.Context) {
	var catalogReq model.CatalogRequest
	if err := c.ShouldBindJSON(&catalogReq); err != nil {
		res.NewResponse(c, 400, "Invalid input", "Please provide valid catalog details")
		return
	}

	userId := c.MustGet("user_id").(int)
	var vendor model.Vendor
	if err := config.DB.Where("id = ?", catalogReq.VendorID).First(&vendor).Error; err != nil {
		res.NewResponse(c, 404, "Vendor not found", nil)
		return
	}
	catalog := model.Catalog{
		Name:        catalogReq.Name,
		Slug:        util.Slugify(catalogReq.Name),
		Description: catalogReq.Description,
		UserID:      userId,
		VendorID:    catalogReq.VendorID,
	}

	if err := config.DB.Create(&catalog).Error; err != nil {
		res.NewResponse(c, 500, "Internal server error", "Could not create catalog")
		return
	}

	if err := config.DB.Where("id = ?", catalog.ID).Preload("User").Preload("Vendor.User").First(&catalog).Error; err != nil {
		res.NewResponse(c, 500, "Internal server error", "Could not retrieve catalog with relations")
		return
	}

	catalogResponse := catalog.ToResponse()
	res.NewResponse(c, 201, "Catalog created successfully", catalogResponse)
	// res.NewResponse(c, 201, "Catalog created successfully", catalog)
}

func GetAllCatalogs(c *gin.Context) {
	var catalogs *[]model.Catalog
	if err := config.DB.Preload("User").Preload("Vendor.User").Find(&catalogs).Error; err != nil {
		res.NewResponse(c, 500, "Failed to retrieve catalogs", nil)
		return
	}

	var catalogResponses []model.CatalogResponse
	for _, catalog := range *catalogs {
		catalogResponses = append(catalogResponses, catalog.ToResponse())
	}

	res.NewResponse(c, 200, "Catalogs retrieved successfully", catalogResponses)
}

func GetCatalogByID(c *gin.Context) {
	catalogID := c.Param("id")
	var catalog model.Catalog

	if err := config.DB.Preload("User").Preload("Vendor.User").First(&catalog, catalogID).Error; err != nil {
		res.NewResponse(c, 404, "Catalog not found", nil)
		return
	}

	catalogResponse := catalog.ToResponse()
	res.NewResponse(c, 200, "Catalog retrieved successfully", catalogResponse)
}

func UpdateCatalog(c *gin.Context) {
	catalogID := c.Param("id")
	var catalogReq model.CatalogRequest

	if err := c.ShouldBindJSON(&catalogReq); err != nil {
		res.NewResponse(c, 400, "Invalid input", "Please provide valid catalog details")
		return
	}

	var catalog model.Catalog
	if err := config.DB.First(&catalog, catalogID).Error; err != nil {
		res.NewResponse(c, 404, "Catalog not found", nil)
		return
	}

	catalog.Name = catalogReq.Name
	catalog.Slug = util.Slugify(catalogReq.Name)
	catalog.Description = catalogReq.Description
	catalog.VendorID = catalogReq.VendorID

	if err := config.DB.Save(&catalog).Error; err != nil {
		res.NewResponse(c, 500, "Internal server error", "Could not update catalog")
		return
	}

	config.DB.Preload("User").Preload("Vendor.User").First(&catalog)

	catalogResponse := catalog.ToResponse()
	res.NewResponse(c, 200, "Catalog updated successfully", catalogResponse)
}

func DeleteCatalog(c *gin.Context) {
	catalogID := c.Param("id")
	var catalog model.Catalog

	if err := config.DB.First(&catalog, catalogID).Error; err != nil {
		res.NewResponse(c, 404, "Catalog not found", nil)
		return
	}

	if err := config.DB.Delete(&catalog).Error; err != nil {
		res.NewResponse(c, 500, "Internal server error", "Could not delete catalog")
		return
	}

	res.NewResponse(c, 200, "Catalog deleted successfully", nil)
}
