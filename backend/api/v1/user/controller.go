package user

import (
	"backend/api/common"
	"backend/api/utils"
	"backend/business"
	"backend/business/user"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service user.Service
}

func NewUserController(service user.Service) *UserController {
	return &UserController{
		service: service,
	}
}

func (uc *UserController) CreateNewUser(c echo.Context) error {

	var body userRequestBody

	if err := c.Bind(&body); err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	_, err := uc.service.CreateNew(body.convertToUserBusiness())
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.SuccessCreated, echo.Map{}))
}

func (uc *UserController) GetCurrentUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaimsUser)
	userId := claims.ID

	userData, err := uc.service.GetCurrent(userId)
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}
	userResponse := convertToUserResponse(userData)
	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"user": userResponse,
	}))

}

func (uc *UserController) Login(c echo.Context) error {
	var body loginRequestBody

	if err := c.Bind(&body); err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	token, err := uc.service.Login(body.Email, body.Password)
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(common.ConstructResponse(business.Success, echo.Map{"token": token}))

}

func (uc *UserController) UpdateUser(c echo.Context) error {
	var body userRequestBody

	if err := c.Bind(&body); err != nil {
		return c.JSON(common.ConstructResponse(business.BadRequest, echo.Map{}))
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaimsUser)
	userId := claims.ID

	userData, err := uc.service.Update(userId, body.Name, body.Address)
	if err != nil {
		return c.JSON(common.ConstructResponse(err.Error(), echo.Map{}))
	}
	userResponse := convertToUserResponse(userData)
	return c.JSON(common.ConstructResponse(business.Success, echo.Map{
		"user": userResponse,
	}))
}
