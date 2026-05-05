package constants

import "net/http"

var (
	ErrBadRequest     = "bad request"
	ErrInternalServer = "internal server error"
	ErrUnauthorized   = "unauthorized"

	ErrFailedToGetUserData    = "failed to get user data"
	ErrFailedToParseUser      = "failed to parse user"
	ErrFailedToParseTokenData = "failed to parse user data"
	ErrTransactionFailed      = "transaction failed"
)

var ()

var ValidationErrorMap = map[string]map[string]string{
	"required": {},
}

type AppError struct {
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(statusCode int, message string) *AppError {
	return &AppError{StatusCode: statusCode, Message: message}
}

var (
	ErrorUnauthorized        = NewAppError(http.StatusUnauthorized, ErrUnauthorized)
	ErrorBadRequest          = NewAppError(http.StatusBadRequest, ErrBadRequest)
	ErrorInternalServer      = NewAppError(http.StatusInternalServerError, ErrInternalServer)
	ErrorFailedToGetUserData = NewAppError(http.StatusInternalServerError, ErrFailedToGetUserData)
	ErrorFailedToParseToken  = NewAppError(http.StatusInternalServerError, ErrFailedToParseTokenData)
	ErrorTransactionFailed   = NewAppError(http.StatusInternalServerError, ErrTransactionFailed)
	ErrorFailedToParseUser   = NewAppError(http.StatusInternalServerError, ErrFailedToParseUser)
)
