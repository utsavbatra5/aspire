package api

import (
	"aspire/api/common"
	"aspire/business"
	"aspire/constants"
	"aspire/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateUserLoan(ctx *gin.Context) {

	var userID int

	// read userID from context that has been set from value from jwt token
	val, ok := ctx.Value(constants.UserID).(string)
	if ok {
		userID, _ = strconv.Atoi(val)
	}

	reqBody := &models.LoanCreateRequest{}

	err := ctx.ShouldBindBodyWith(reqBody, binding.JSON)
	if err != nil {
		common.SendBadRequest(ctx, constants.ErrRequestBodyBindMsg, constants.ErrRequestBodyBindCode, false, err)
		return
	}

	if reqBody.Amount == 0 || reqBody.Freq == 0 {
		common.SendBadRequest(ctx, constants.ErrInvalidPayload, constants.ErrRequestBodyValidationCode, false, err)
		return
	}

	response, err := business.CreateUserLoan(ctx, *reqBody, userID)
	if err != nil {
		common.SendInternalServerError(ctx, constants.ErrCommonMsg, constants.ErrExternalServiceFailureCode, false, err)
		return
	}

	common.SendStatusOK(ctx, false, response)
}

func GetUserLoans(ctx *gin.Context) {
	var userID int

	// read userID from context that has been set from value from jwt token
	val, ok := ctx.Value(constants.UserID).(string)
	if ok {
		userID, _ = strconv.Atoi(val)
	}

	response, err := business.GetUserLoans(ctx, userID)
	if err != nil {
		common.SendInternalServerError(ctx, constants.ErrCommonMsg, constants.ErrExternalServiceFailureCode, false, err)
		return
	}

	common.SendStatusOK(ctx, false, response)
}

func ApproveLoan(ctx *gin.Context) {
	req := &models.ApproveLoan{}
	err := ctx.BindQuery(req)
	if err != nil {
		common.SendBadRequest(ctx, constants.ErrRequestBodyBindMsg, constants.ErrRequestBodyBindCode, false, err)
		return
	}

	if req.Loanid == "" {
		common.SendBadRequest(ctx, constants.ErrInvalidPayload, constants.ErrRequestBodyValidationCode, false, err)
		return
	}

	response, err := business.ApproveLoan(ctx, req.Loanid)
	if err != nil {
		common.SendInternalServerError(ctx, constants.ErrCommonMsg, constants.ErrExternalServiceFailureCode, false, err)
		return
	}

	common.SendStatusOK(ctx, false, response)
}

func PayEMI(ctx *gin.Context) {
	reqBody := &models.PayEMI{}
	err := ctx.ShouldBindWith(reqBody, binding.JSON)

	if err != nil {
		common.SendBadRequest(ctx, constants.ErrRequestBodyBindMsg, constants.ErrRequestBodyBindCode, false, err)
		return
	}

	if reqBody.Loanid == "" || reqBody.Emiamount == 0 {
		common.SendBadRequest(ctx, constants.ErrInvalidPayload, constants.ErrRequestBodyValidationCode, false, err)
		return
	}

	response, err := business.PayEMI(ctx, reqBody.Loanid, reqBody.Emiamount)
	if err != nil {
		common.SendInternalServerError(ctx, constants.ErrCommonMsg, constants.ErrExternalServiceFailureCode, false, err)
		return
	}

	common.SendStatusOK(ctx, false, response)
}
