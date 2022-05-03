package middleware

import (
	"aspire/api/common"
	"aspire/config"
	"aspire/constants"
	"aspire/models"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthenticationMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var accessToken string

		token := ctx.Request.Header.Get("Authorization")
		tokenString := strings.Split(token, " ")
		if len(tokenString) == 2 && tokenString[0] == "Bearer" {
			accessToken = tokenString[1]
		} else {
			common.SendBadRequest(ctx, constants.HttpErrMessageBadRequest, constants.ErrRequestTokenValidationCode, false, errors.New("Empty AccessToken"))
			ctx.Abort()
			return
		}

		accessTokenData, err := decodeUserToken(accessToken)
		if err != nil {
			common.SendUnauthorized(ctx, constants.ErrUnauthorizedMsg, constants.ErrUnauthorizedCode, false, err)
			ctx.Abort()
			return
		} else {
			ctx.Set(constants.UserID, accessTokenData.UserID)
			ctx.Next()
		}
	}
}

func decodeUserToken(tokenID string) (models.TokenUserData, error) {

	tokenData := models.TokenUserData{}

	jwtKey := config.Conf.JwtKey

	token, err := jwt.Parse(tokenID, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		return models.TokenUserData{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenData.Expiration = (claims["Expiration"]).(string)
		tokenData.IssuedAt = (claims["IssuedAt"]).(string)
		tokenData.Issuer = (claims["Issuer"]).(string)
		tokenData.UserID = (claims["UserID"]).(string)
	} else {
		return tokenData, fmt.Errorf("invalid error")
	}

	return tokenData, nil
}
