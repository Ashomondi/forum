package auth

import "golang.org/x/crypto/bcrypt"

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) Register(user User) error {
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
		return User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return User{}, err
	}
	return user, nil
}
