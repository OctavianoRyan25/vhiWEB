package model

import "time"

type Vendor struct {
	ID        int
	Name      string
	Company   string
	Email     string
	Phone     string
	Address   string
	UserID    int
	User      *User
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VendorRequest struct {
	Name    string `json:"name" validate:"required"`
	Company string `json:"company" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type VendorResponse struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Company   string       `json:"company"`
	Email     string       `json:"email"`
	Phone     string       `json:"phone"`
	Address   string       `json:"address"`
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `json:"created_at"`
}

func (v *Vendor) ToResponse() *VendorResponse {
	return &VendorResponse{
		ID:        v.ID,
		Name:      v.Name,
		Company:   v.Company,
		Email:     v.Email,
		Phone:     v.Phone,
		Address:   v.Address,
		User:      *(v.User.ToResponse()),
		CreatedAt: v.CreatedAt,
	}
}
