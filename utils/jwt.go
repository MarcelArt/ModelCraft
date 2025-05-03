package utils

import (
	"strconv"
	"time"

	"github.com/MarcelArt/ModelCraft/config"
	"github.com/golang-jwt/jwt/v5"
)

type ILoginable interface {
	AccessClaims(exp int64) jwt.MapClaims
	RefreshClaims(isRememeber bool) jwt.MapClaims
}

func GenerateTokenPair(user ILoginable, isRemember bool) (string, string, error) {
	accessToken, err := generateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateRefreshToken(user, isRemember)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func generateAccessToken(user ILoginable) (string, error) {
	claims := user.AccessClaims(time.Now().Add(time.Minute * 5).Unix())

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	t, err := token.SignedString([]byte(config.Env.JwtSecret))

	return t, err
}

func generateRefreshToken(user ILoginable, isRemember bool) (string, error) {
	claims := user.RefreshClaims(isRemember)

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	t, err := token.SignedString([]byte(config.Env.JwtSecret))

	return t, err
}

func ParseToken(t string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Env.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func ClaimsNumberToString(i interface{}) string {
	number := i.(float64)

	return strconv.Itoa(int(number))
}
