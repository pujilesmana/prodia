package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	middleware.ErrJWTInvalid.Code = 401
	middleware.ErrJWTInvalid.Message = "unauthorized"
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "unauthorized"
}

type JwtCustomClaimsUser struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
