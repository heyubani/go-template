package service

import (
	appErrors "github.com/heyubani/go-template/core/app-errors"
	"github.com/heyubani/go-template/interfaces"
	models "github.com/heyubani/go-template/modules/auth/model"
	"github.com/heyubani/go-template/modules/auth/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService struct {
	logger     interfaces.ILogger
	repository repository.AuthRepository
}

func NewAuthService(logger interfaces.ILogger) AuthService {
	return AuthService{
		logger:     logger,
		repository: repository.NewAuthRepository(logger),
	}
}

func (as *AuthService) GetUser(id primitive.ObjectID) (models.User, interfaces.IAppError) {
	user, err := as.repository.GetByID(id)

	if err != nil {
		as.logger.Errorf("Error fetching user", err)
		return models.User{}, appErrors.InternalServerError
	}

	return user, nil
}
