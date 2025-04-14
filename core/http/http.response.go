package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heyubani/go-template/interfaces"
)

type statusType string

var (
	failed         statusType = "failed"
	success        statusType = "success"
	successMessage string     = "Request was successfull"
	errorMessage   string     = "Returned with an error , kindly check the error field"
)

type Response struct {
	Status  statusType `json:"status,omitempty"`
	Message string     `json:"message,omitempty"`
	Data    any        `json:"data,omitempty"`
	Error   any        `json:"error,omitempty"`
}

func response(context *gin.Context, data any, status statusType, message string, code int) {
	var response Response

	err, isError := data.(interfaces.IAppError)
	if isError {
		response = Response{
			Status:  status,
			Message: message,
			Error:   err.Error(),
		}

	} else {
		response = Response{
			Data:    data,
			Status:  success,
			Message: message,
		}
	}
	context.Header("Allow-Control-Allow-Origin", "*")
	context.JSON(code, response)

}

func JsonOk(context *gin.Context, data interface{}) {
	response(context, data, success, successMessage, http.StatusOK)
}

func JsonOtpOk(context *gin.Context, verify bool) {
	var msg string = "Kindly check your email for otp"
	if verify {
		msg = "Kindly check your email for the next step to reset password"
	}
	response(context, nil, success, msg, http.StatusOK)
}

func JsonCreated(context *gin.Context, data any, module string) {
	response(context, data, success, fmt.Sprintf("%v created successfully", module), http.StatusCreated)
}

func JsonDelete(context *gin.Context, module string) {
	response(context, nil, success, fmt.Sprintf("%v Deleted Successfully ", module), http.StatusOK)
}

func JsonModified(context *gin.Context, data any, module string) {
	response(context, data, success, fmt.Sprintf("%v was modified successfully", module), http.StatusAccepted)
}

func JsonError(context *gin.Context, err interfaces.IAppError) {
	response(context, err, failed, errorMessage, err.GetCode())
}
