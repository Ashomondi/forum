package auth

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func validateUser(user User) error {
	if strings.TrimSpace(user.Username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email is required")
	}
	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}

func (s *Service) Register(user User) error {
	if err := validateUser(user); err != nil {
		return err
	}

	_, err := s.Repo.GetUserByEmail(user.Email)
	if err == nil {
		// No error means the user was found so by default, email already taken.
		return errors.New("email already exists")
	}
	if !errors.Is(err, errors.New("user not found")) {
		// Something went wrong with the DB , so don't proceed.
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.Repo.CreateUser(user)
}

func (s *Service) Login(email, password string) (User, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		// Return a generic message, don't reveal whether the email exists.
		return User{}, errors.New("invalid email or password")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return User{}, errors.New("invalid email or password")
	}

	return user, nil
}
