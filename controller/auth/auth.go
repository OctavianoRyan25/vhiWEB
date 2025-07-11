package auth

import (
	"github.com/OctavianoRyan25/VhiWEB/config"
	"github.com/OctavianoRyan25/VhiWEB/model"
	"github.com/OctavianoRyan25/VhiWEB/res"
	"github.com/OctavianoRyan25/VhiWEB/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Register(c *gin.Context) {
	var userReq model.UserRegisterRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		res.NewResponse(c, 400, "Invalid input", "Please provide valid user details")
		return
	}

	validate := validator.New()
	if err := validate.Struct(&userReq); err != nil {
		errors := []string{}
		for _, error := range err.(validator.ValidationErrors) {
			// Format the error message to be more user-friendly
			errors = append(errors, error.Field()+" is "+error.Error())
		}
		res.NewResponse(c, 400, "Validation error", errors)
		return
	}

	var userExist model.User

	config.DB.Where("email = ?", userReq.Email).First(&userExist)

	if userExist.ID != 0 {
		res.NewResponse(c, 409, "Conflict", "Email already registered")
		return
	}

	hashedPassword, err := util.Hash(userReq.Password)
	if err != nil {
		res.NewResponse(c, 500, "Internal server error", "Could not hash password")
		return
	}

	user := model.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: hashedPassword,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		res.NewResponse(c, 500, "Internal server error", "Could not register user")
		return
	}

	res.NewResponse(c, 201, "Registration successful", nil)
}

func Login(c *gin.Context) {
	var userReq model.UserLoginRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		res.NewResponse(c, 400, "Invalid input", "Please provide valid login details")
		return
	}

	validate := validator.New()
	if err := validate.Struct(&userReq); err != nil {
		res.NewResponse(c, 400, "Validation error", err.Error())
		return
	}

	var user model.User

	config.DB.Where("email = ?", userReq.Email).First(&user)

	if user.ID == 0 {
		res.NewResponse(c, 401, "Unauthorized", "Invalid email or password")
		return
	}

	ok := util.CompareHashPassword(user.Password, userReq.Password)

	if !ok {
		res.NewResponse(c, 401, "Unauthorized", "Invalid email or password")
		return
	}

	token, err := util.GenerateJWT(user.ID, user.Email)
	if err != nil {
		res.NewResponse(c, 500, "Internal server error", "Could not generate token")
		return
	}

	res.NewResponse(c, 200, "Login successful", gin.H{
		"token": token,
	})
}
