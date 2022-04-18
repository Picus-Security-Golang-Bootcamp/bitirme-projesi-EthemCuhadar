package helper

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type JwtMetaData struct {
	UserID  string
	IsAdmin bool
}

// TokenValid checks validation of token
func TokenValid(r *http.Request, cfg *config.Config) error {
	token, err := VerifyToken(r, cfg)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// VerifyToken will verify token
func VerifyToken(r *http.Request, cfg *config.Config) (*jwt.Token, error) {

	// Extract
	tokenString := ExtractToken(r)

	// Parse
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(cfg.JWTConfig.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken will got token from request
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// ExtractMetaData will get information in the token
func ExtractMetaData(r *http.Request, cfg *config.Config) (*JwtMetaData, error) {

	// Verify
	token, err := VerifyToken(r, cfg)
	if err != nil {
		return nil, err
	}

	// Data
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		isAdmin, ok := claims["isAdmin"].(bool)
		if !ok {
			zap.L().Debug("ExtractMetaData", zap.Reflect("err 1:", err))
			return nil, err
		}
		userId, ok := claims["userId"].(string)
		if !ok {
			zap.L().Debug("ExtractMetaData", zap.Reflect("err 2:", err))
			return nil, err
		}
		return &JwtMetaData{
			IsAdmin: isAdmin,
			UserID:  userId,
		}, nil
	}
	return nil, err
}
