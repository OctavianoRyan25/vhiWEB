package controller

import (
	"github.com/OctavianoRyan25/VhiWEB/config"
	"github.com/OctavianoRyan25/VhiWEB/model"
	"github.com/OctavianoRyan25/VhiWEB/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateVendor(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	var vendorReq model.VendorRequest

	if err := c.ShouldBindJSON(&vendorReq); err != nil {
		res.NewResponse(c, 400, "Invalid input", "Please provide valid vendor details")
		return
	}

	validate := validator.New()
	if err := validate.Struct(&vendorReq); err != nil {
		errors := []string{}
		for _, error := range err.(validator.ValidationErrors) {
			errors = append(errors, error.Field()+" is "+error.Error())
		}
		res.NewResponse(c, 400, "Validation error", errors)
		return
	}

	vendor := model.Vendor{
		Name:    vendorReq.Name,
		Company: vendorReq.Company,
		Email:   vendorReq.Email,
		Phone:   vendorReq.Phone,
		Address: vendorReq.Address,
		UserID:  userId,
	}

	if err := config.DB.Create(&vendor).Error; err != nil {
		res.NewResponse(c, 500, "Internal server error", "Could not create vendor")
		return
	}

	config.DB.Where("id = ?", vendor.ID).Preload("User").First(&vendor)

	vendorResponse := vendor.ToResponse()

	res.NewResponse(c, 201, "Vendor created successfully", vendorResponse)
}
