package auth

import (
	"errors"
	"github.com/VATUSA/api-v3/internal/config"
	"github.com/VATUSA/api-v3/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

func GetRequestUserID(c echo.Context) (*uint64, error) {
	token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return nil, errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return nil, errors.New("failed to cast claims as jwt.MapClaims")
	}
	subject, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}
	userId, err := strconv.ParseUint(subject, 10, 64)
	if err != nil {
		return nil, err
	}
	return &userId, nil
}

type Claims struct {
	CID uint64 `json:"cid"`
	jwt.RegisteredClaims
}

func CreateJWTForController(controller *database.Controller) (*string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		CID: 0,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			NotBefore: nil,
			IssuedAt:  nil,
			ID:        "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(config.JWTSecretKey)
	if err != nil {
		return nil, err
	}
	return &tokenString, err
}

func GetControllerForJWT(tokenString string) (*database.Controller, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	cid := claims.CID
	controller, err := database.FetchControllerByCID(cid)
	if err != nil {
		return nil, err
	}
	return controller, nil
}
