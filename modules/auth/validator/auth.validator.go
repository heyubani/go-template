package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/heyubani/go-template/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// i'm writing this call so you know what i expect for validator

type AuthValidator struct {
	Logger    interfaces.ILogger
	Validator validator.Validate
}

func NewAuthValidator(logger interfaces.ILogger) AuthValidator {
	return AuthValidator{
		Logger:    logger,
		Validator: *validator.New(),
	}
}

func (av *AuthValidator) ValidateID(id primitive.ObjectID) interfaces.IAppError {

	// this is how to use the validator class

	// if err := av.Validator.Struct(); err != nil {
	// 	return appErrors.BadRequestError(err.Error())
	// }

	// you can call the repository to fetch data from db

	return nil
}
