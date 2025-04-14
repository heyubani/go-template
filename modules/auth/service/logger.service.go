package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	appErrors "github.com/heyubani/go-template/core/app-errors"
	"github.com/heyubani/go-template/interfaces"
	models "github.com/heyubani/go-template/modules/auth/model"
	"github.com/heyubani/go-template/modules/auth/repository"
)

type LoggerService struct {
	logger     interfaces.ILogger
	repository repository.AuthRepository
}

func NewLoggerService(logger interfaces.ILogger) AuthService {
	return AuthService{
		logger:     logger,
		repository: repository.NewAuthRepository(logger),
	}
}

func (as *AuthService) GetLogger(id primitive.ObjectID) (models.User, interfaces.IAppError) {
	user, err := as.repository.GetByID(id)

	if err != nil {
		as.logger.Errorf("Error fetching user", err)
		return models.User{}, appErrors.InternalServerError
	}

	return user, nil
}
