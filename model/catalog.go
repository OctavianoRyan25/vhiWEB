package model

import "time"

type Catalog struct {
	ID          int
	Name        string
	Slug        string
	Description string
	UserID      int
	User        *User
	VendorID    int
	Vendor      *Vendor
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CatalogRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	VendorID    int    `json:"vendor_id" validate:"required"`
}

type CatalogResponse struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	Description string         `json:"description"`
	User        UserResponse   `json:"user"`
	Vendor      VendorResponse `json:"vendor"`
	CreatedAt   time.Time      `json:"created_at"`
}

func (c *Catalog) ToResponse() CatalogResponse {
	var userResp UserResponse
	if c.User != nil {
		userResp = *c.User.ToResponse()
	}

	var vendorResp VendorResponse
	if c.Vendor != nil {
		vendorResp = *c.Vendor.ToResponse()
	}

	return CatalogResponse{
		ID:          c.ID,
		Name:        c.Name,
		Slug:        c.Slug,
		Description: c.Description,
		User:        userResp,
		Vendor:      vendorResp,
		CreatedAt:   c.CreatedAt,
	}
}
