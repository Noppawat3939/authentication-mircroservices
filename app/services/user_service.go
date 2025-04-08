package services

import (
	"auth-microservice/models"
	"auth-microservice/repository"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) RegisterUser(user *models.User) (*models.User, error) {
	exitingUser, err := s.userRepo.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if exitingUser != nil {
		return nil, errors.New("email alreay exits")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("failed hashing password", err)
		return nil, err
	}

	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)

	err = s.userRepo.Insert(user)

	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}

func (s *UserService) LoginUser(user *models.User) (*models.User, error) {
	foundUser, err := s.userRepo.FindByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	if foundUser != nil {
		return nil, errors.New("email or password invalid")
	}

	passwordIsValid, err := verifyPassword(user.Password, foundUser.Password)

	if err != nil {
		return nil, errors.New("email or password invalid")
	}

	if !passwordIsValid {
		return nil, errors.New("email or password invalid")
	}

	return foundUser, nil
}

func verifyPassword(userPassword string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))

	if err != nil {
		return false, err
	}

	return true, nil
}
