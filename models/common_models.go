package models

import (
	"Backend/types"
	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	jwt.StandardClaims
	Login string
	Role  types.Role
}
