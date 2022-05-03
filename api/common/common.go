package common

import (
	"aspire/logger"
	"aspire/models"
	"net/http"

	"github.com/angel-one/bbe-discovery/constants"
	"github.com/gin-gonic/gin"
)

func getData(isJSONArray bool, data interface{}) interface{} {
	if data != nil {
		return data
	}

	type emptyData string

	if isJSONArray {
		return []emptyData{}
	}
	return ""
}

func SendBadRequest(ctx *gin.Context, msg, code string, isJSONArray bool, err error) {
	statusCode := http.StatusBadRequest
	response := buildErrorResponse(msg, code, isJSONArray, nil)
	logger.ZeroLog().Error().Err(err).Msg("bad request")
	ctx.JSON(statusCode, response)
}

func SendNotFound(ctx *gin.Context, msg, code string, isJSONArray bool, err error) {
	statusCode := http.StatusNotFound
	response := buildErrorResponse(msg, code, isJSONArray, nil)
	logger.ZeroLog().Error().Err(err).Msg("not found")
	ctx.JSON(statusCode, response)
}

func SendInternalServerError(ctx *gin.Context, msg, code string, isJSONArray bool, err error) {
	statusCode := http.StatusInternalServerError
	response := buildErrorResponse(msg, code, isJSONArray, nil)
	logger.ZeroLog().Error().Err(err).Msg("internal server error")
	ctx.JSON(statusCode, response)
}

func SendSomethingWrong(ctx *gin.Context, msg, code string, isJSONArray bool, err error) {
	statusCode := http.StatusPreconditionFailed
	response := buildErrorResponse(msg, code, isJSONArray, nil)
	logger.ZeroLog().Error().Err(err).Msg("something went wrong")
	ctx.JSON(statusCode, response)
}

func SendInvalidServerError(ctx *gin.Context, msg, code string, isJSONArray bool, err error) {
	statusCode := http.StatusInternalServerError
	response := buildErrorResponse(msg, code, isJSONArray, nil)
	logger.ZeroLog().Error().Err(err).Msg("invalid server error")
	ctx.JSON(statusCode, response)
}

func SendStatusOK(ctx *gin.Context, isJSONArray bool, data interface{}) {
	responseData := getData(isJSONArray, data)
	ctx.Set(constants.LogResponseKey, responseData)
	ctx.JSON(http.StatusOK, models.APIResponse{
		Status: constants.APIRespSuccessKey,
		Data:   responseData,
	})
}

func SendUnauthorized(ctx *gin.Context, msg, code string, isJSONArray bool, err error) {
	response := buildErrorResponse(msg, code, isJSONArray, nil)
	logger.ZeroLog().Error().Err(err).Msg("unauthorized user")
	ctx.JSON(http.StatusUnauthorized, response)
}

//Token expired, user should relogin
func SendTokenExpired(ctx *gin.Context, msg, code string, isJSONArray bool, err error) {
	response := buildErrorResponse(msg, code, isJSONArray, nil)
	logger.ZeroLog().Error().Err(err).Msg("token expired")
	ctx.JSON(constants.StatusTokenExpired, response)
}

func buildErrorResponse(msg string, code string, isJSONArray bool, data interface{}) models.APIResponse {
	return models.APIResponse{
		Status:    constants.APIRespErrorKey,
		Message:   msg,
		ErrorCode: code,
		Data:      getData(isJSONArray, data),
	}
}
