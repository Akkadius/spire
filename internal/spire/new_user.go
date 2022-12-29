package spire

import (
	"errors"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/models"
	"github.com/sirupsen/logrus"
	"time"
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

func (s UserService) CreateUser(username string, password string) (models.User, error) {
	// check if user exists
	var users []models.User
	s.db.GetSpireDb().Where("user_name = ?", username).Find(&users)
	if len(users) > 0 {
		return models.User{}, errors.New("User already exists")
	}

	// hash password
	hash, err := s.crypt.GeneratePassword(password)
	if err != nil {
		return models.User{}, err
	}

	// create
	var newUser models.User
	s.db.GetSpireDb().FirstOrCreate(
		&newUser, models.User{
			UserName:  username,
			FullName:  username,
			Password:  hash,
			Provider:  "local",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	)

	if newUser.ID > 0 {
		s.logger.Infof("[user] Created user ID [%v]", newUser.ID)
	}

	return newUser, nil
}
