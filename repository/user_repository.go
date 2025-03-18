package repository

import (
	db "auth-microservice/database"
	"auth-microservice/internal/utils"
	"auth-microservice/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	collection := db.GetCollection("auth-go", "users")

	return &UserRepository{collection: collection}
}

func (r *UserRepository) Insert(user *models.User) error {
	ctx, cancel := utils.CreateContextWithTimeout()
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	ctx, cancel := utils.CreateContextWithTimeout()
	defer cancel()

	var user models.User

	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return &user, nil
}
