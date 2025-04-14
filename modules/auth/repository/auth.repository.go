package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/heyubani/go-template/database"
	"github.com/heyubani/go-template/interfaces"
	models "github.com/heyubani/go-template/modules/auth/model"
)

var (
	dBName       = "database_name" // database name
	dbCollection = "auth"          // collection name
	ctx          context.Context
)

type AuthRepository struct {
	Logger     interfaces.ILogger
	Collection *mongo.Collection
}

func NewAuthRepository(logger interfaces.ILogger) AuthRepository {
	collection := database.GetMongoCollection(dBName, dbCollection)
	return AuthRepository{
		Logger:     logger,
		Collection: collection,
	}
}

func (ar *AuthRepository) GetByID(id primitive.ObjectID) (models.User, error) {
	var user models.User
	filter := primitive.M{
		"_id": primitive.M{"$in": id},
	}

	err := ar.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		ar.Logger.Errorf("Error finding products", err)
		return user, err
	}

	return user, nil
}
