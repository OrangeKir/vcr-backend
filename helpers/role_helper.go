package helpers

import (
	"Backend/types"
	"bytes"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"
)

const authServerUrl = "http://localhost:5080/validate-token"

func ValidateToken(tokenString string) (bool, string, types.Role) {
	request := Request{Token: tokenString}

	jsonBody, err := json.Marshal(request)

	httpRequest, err := http.NewRequest("GET", authServerUrl, bytes.NewBuffer(jsonBody))

	client := &http.Client{}
	resp, err := client.Do(httpRequest)

	if err != nil {
		return false, "", ""
	}

	var response Response
	json.NewDecoder(resp.Body).Decode(&response)

	if err != nil || !response.IsValid {
		return false, "", ""
	}

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, TokenModel{})
	claims, _ := token.Claims.(*TokenModel)

	if err != nil {
		return false, "", ""
	}

	return true, claims.Login, claims.Role
}

type Response struct {
	IsValid bool
}

type Request struct {
	Token string
}

type TokenModel struct {
	jwt.StandardClaims
	Login string
	Role  types.Role
}
