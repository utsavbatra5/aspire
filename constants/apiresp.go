package constants

//API Response Status and Error Codes
const (
	APIRespSuccessKey             = "success"
	APIRespErrorKey               = "error"
	ErrRequestBodyBindCode        = "RB-101"
	ErrRequestBodyValidationCode  = "RB-102"
	ErrExternalServiceFailureCode = "SYS-101"
	ErrDatabaseFailureCode        = "SYS-102"
	ErrNoDataFoundCode            = "SYS-103"
	ErrUnauthorizedCode           = "UA-101"
	ErrTokenExpiredCode           = "UA-102"
	ErrNoAppVersionCode           = "AV-101"

	ErrRequestTokenValidationCode = "RB-103"

	ErrRequestBodyBindMsg        = "request body bind error"
	ErrRequestBodyValidationMsg  = "request body validation error"
	ErrExternalServiceFailureMsg = "external service failure error"
	ErrDatabaseFailureMsg        = "database failure error"
	ErrCommonMsg                 = "something went wrong"
	ErrNoDataFoundMsg            = "data not found"
	ErrUnauthorizedMsg           = "unauthorized token"
	ErrInvalidHashMsg            = "invalid hash"
	ErrTokenExpired              = "token expired"
	ErrInvalidPayload            = "Invalid payload"
	ErrNoSuchUser                = "No such user"
	ErrNoAppVersion              = "No valid app version"

	HttpErrMessageBadRequest = "Bad Request"

	Success = "Success"

	StatusOKCode       = "201"
	StatusCreatedCode  = "201"
	StatusTokenExpired = 440
)
