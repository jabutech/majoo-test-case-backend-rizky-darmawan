package service

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/majoo-test/models/domain"
	"github.com/majoo-test/models/web"
	"github.com/majoo-test/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(req web.LoginRequest) (domain.User, error)
	GenerateToken(user domain.User) (string, error)
}

type service struct {
	repository repository.Repository
}

func NewServiceUser(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) Login(req web.LoginRequest) (domain.User, error) {
	// Get payload
	username := req.UserName
	password := req.Password

	// Find user by username
	user, err := s.repository.FindByUserName(username)
	if err != nil {
		return user, err
	}

	// If user not found
	if user.ID == 0 {
		return user, errors.New("email or password incorrect")
	}

	// If user is available, compare password hash with password from request use bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("email or password incorrect")
	}

	return user, nil
}

type Claim struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func (s *service) GenerateToken(user domain.User) (string, error) {
	// Create 1 day
	expirationTime := time.Now().AddDate(0, 0, 1)

	// Create clain for payload token
	claim := Claim{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// load env
	godotenv.Load("../.env")

	SecretJWT := os.Getenv("SECRET_KEY")
	// Signed token with secret key
	signedToken, err := token.SignedString([]byte(SecretJWT))
	if err != nil {
		return signedToken, err
	}

	// If success, return token
	return signedToken, nil
}
