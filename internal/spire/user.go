package spire

import (
	"errors"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/models"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	LOGIN_PROVIDER_LOCAL  = "local"
	LOGIN_PROVIDER_GITHUB = "github"
)

type UserService struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
	crypt  *encryption.Encrypter
}

func NewUserService(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	crypt *encryption.Encrypter,
) *UserService {
	return &UserService{
		db:     db,
		logger: logger,
		crypt:  crypt,
	}
}

func (s UserService) CreateUser(user models.User) (models.User, error) {
	// check if user exists
	var users []models.User
	s.db.GetSpireDb().
		Where("user_name = ? and provider = ?", user.UserName, user.Provider).
		Find(&users)
	if len(users) > 0 {
		return models.User{}, errors.New("User already exists")
	}

	if len(user.UserName) < 3 {
		return models.User{}, errors.New("Username must be at least 3 characters")
	}

	if len(user.Password) < 8 && user.Provider == LOGIN_PROVIDER_LOCAL {
		return models.User{}, errors.New("Password must be at least 8 characters")
	}

	if user.Provider == LOGIN_PROVIDER_LOCAL {
		hash, err := s.crypt.GeneratePassword(user.Password)
		if err != nil {
			return models.User{}, err
		}
		user.Password = hash
	}

	// defaults
	user.CreatedAt = time.Time{}
	user.UpdatedAt = time.Time{}

	// create
	var newUser models.User
	s.db.GetSpireDb().FirstOrCreate(&newUser, user)

	return newUser, nil
}

func (s UserService) CheckUserLogin(username string, password string) (bool, error) {
	var user models.User
	s.db.GetSpireDb().Where("user_name = ? and provider = ?", username, LOGIN_PROVIDER_LOCAL).First(&user)

	if user.ID == 0 {
		return false, errors.New("User does not exist")
	}

	match, err := s.crypt.ComparePassword(password, user.Password)
	if err != nil {
		return false, err
	}

	return match, nil
}
