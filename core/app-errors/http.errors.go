package appErrors

import (
	"net/http"

	"github.com/heyubani/go-template/interfaces"
)

type LabelText string

var (
	InternalErrorLabel                LabelText = "internal-error"
	UnsupportedContentTypeLabel       LabelText = "unsupported-content-type"
	LargeRequestLabel                 LabelText = "large-request"
	UnavailableJsonContentHeaderLabel LabelText = "unavailable-json-header"
	DataNotFoundLabel                 LabelText = "data-not-found"
	DatabaseErrorLabel                LabelText = "database-server-error"
	BadRequestLabel                   LabelText = "bad-request"
	UnauthorizedErrorLabel            LabelText = "unauthorized-request"
)

type httpError struct {
	code    int
	message string

	//This will be used to identify some segments of error we will have so as to enable proper logging
	label LabelText
}

// Sentinel Errors
var (
	InternalServerError          = newHttpError(http.StatusInternalServerError, "Internal server error", InternalErrorLabel)
	UnavailableJsonContentType   = newHttpError(http.StatusBadRequest, "Request lacks application/json content-type header", UnavailableJsonContentHeaderLabel)
	UnsupportedContentTypeHeader = newHttpError(http.StatusUnprocessableEntity, "Unsupported content type header", UnsupportedContentTypeLabel)
	RequestBodyTooLarge          = newHttpError(http.StatusRequestEntityTooLarge, "Request body too large", LargeRequestLabel)
	DataBaseServerError          = newHttpError(http.StatusInternalServerError, "Internal data server error", DatabaseErrorLabel)
)

func (err *httpError) Error() string {
	return err.message
}
func (err *httpError) GetCode() int {
	return err.code
}

func (err *httpError) GetLabel() string {
	return string(err.label)
}

func newHttpError(code int, message string, label LabelText) interfaces.IAppError {
	return &httpError{
		code:    code,
		message: message,
		label:   label,
	}
}

func BadRequestError(message string) interfaces.IAppError {
	return newHttpError(http.StatusBadRequest, message, BadRequestLabel)
}

func PreconditionFailureError(message string) interfaces.IAppError {
	return newHttpError(http.StatusPreconditionRequired, message, BadRequestLabel)
}

func ResourceNotFoundError(message string) interfaces.IAppError {
	return newHttpError(http.StatusNotFound, message, DataNotFoundLabel)
}

func InternalErrorMsg(message string) interfaces.IAppError {
	return newHttpError(http.StatusInternalServerError, message, InternalErrorLabel)
}

func UnauthorizedError(message string) interfaces.IAppError {
	return newHttpError(http.StatusUnauthorized, message, UnauthorizedErrorLabel)
}

func DataNotFoundError(message string) interfaces.IAppError {
	return newHttpError(http.StatusNotFound, message, DataNotFoundLabel)
}

func ForbiddenError(message string) interfaces.IAppError {
	return newHttpError(http.StatusForbidden, message, DataNotFoundLabel)
}

func PaymentRequiredError(message string) interfaces.IAppError {
	return newHttpError(http.StatusPaymentRequired, message, DataNotFoundLabel)
}
