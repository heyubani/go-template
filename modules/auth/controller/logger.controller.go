package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	appErrors "github.com/heyubani/go-template/core/app-errors"
	response "github.com/heyubani/go-template/core/http"
	"github.com/heyubani/go-template/interfaces"
	"github.com/heyubani/go-template/modules/auth/service"
)

// this a sample controller for you to follow as a guide for others

type LoggerController struct {
	logger  interfaces.ILogger
	service service.AuthService
}

// sample boiler plate
func NewLoggerController(logger interfaces.ILogger) *LoggerController {
	return &LoggerController{
		logger:  logger,
		service: service.NewLoggerService(logger),
	}
}

func (ac *LoggerController) RecordDetails(context *gin.Context) {
	id, exist := context.Params.Get("id")
	if !exist {
		response.JsonError(context, appErrors.BadRequestError("Invalid user id"))
		return
	}

	if id == "" {
		response.JsonError(context, appErrors.BadRequestError("Invalid user id"))
		return
	}

	objID, errS := primitive.ObjectIDFromHex(id)

	if errS != nil {
		response.JsonError(context, appErrors.BadRequestError("Invalid user id"))
		return
	}

	user, err := ac.service.GetUser(objID)

	if err != nil {
		response.JsonError(context, err)
		return
	}

	response.JsonOk(context, user)
}
