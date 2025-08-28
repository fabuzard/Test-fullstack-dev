package service

import (
	"inventory_backend/config"
	"inventory_backend/dto"
	"inventory_backend/model"
	"inventory_backend/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetByID(id int) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(id int, user model.User) (model.User, error)
	Delete(id int) error
	Login(email, password string) (dto.AuthResponse, error)
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
func (s *userService) GetByID(id int) (model.User, error) {
	return s.repo.GetByID(id)
}

// register user
func (s *userService) Create(user model.User) (model.User, error) {
	existingUser, _ := s.repo.GetByEmail(user.Email)
	if existingUser.ID != 0 {
		return model.User{}, nil // User already exists
	}
	// Hash password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	user.Password = string(hashedPassword)
	return s.repo.Create(user)
}
func (s *userService) Update(id int, user model.User) (model.User, error) {
	return s.repo.Update(id, user)
}
func (s *userService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *userService) Login(email, password string) (dto.AuthResponse, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return dto.AuthResponse{}, err
	}
	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return dto.AuthResponse{}, err
	}
	// Generate JWT token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetJWTSecret()))
	if err != nil {
		return dto.AuthResponse{}, err
	}

	response := dto.AuthResponse{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email,
		Token:    tokenString,
	}

	return response, nil
}
