package spire

import (
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/models"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	LoginProviderLocal  = "local"
	LoginProviderGithub = "github"
)

type UserService struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
	crypt  *encryption.Encrypter
	cache  *gocache.Cache
}

func NewUserService(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	crypt *encryption.Encrypter,
	cache *gocache.Cache,
) *UserService {
	return &UserService{
		db:     db,
		logger: logger,
		crypt:  crypt,
		cache:  cache,
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

	if len(user.Password) < 8 && user.Provider == LoginProviderLocal {
		return models.User{}, errors.New("Password must be at least 8 characters")
	}

	if user.Provider == LoginProviderLocal {
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

func (s UserService) CheckUserLogin(username string, password string) (bool, error, models.User) {
	var user models.User
	s.db.GetSpireDb().Where("user_name = ? and provider = ?", username, LoginProviderLocal).First(&user)

	if user.ID == 0 {
		return false, errors.New("User does not exist"), models.User{}
	}

	match, err := s.crypt.ComparePassword(password, user.Password)
	if err != nil {
		return false, err, models.User{}
	}

	return match, nil, user
}

func (s UserService) PurgeUserCache(userId uint) {
	s.cache.Delete(fmt.Sprintf("active-connection-%v", userId))
	s.cache.Delete(fmt.Sprintf("active-connection-%v-default", userId))
	s.cache.Delete(fmt.Sprintf("active-connection-%v-eqemu_content", userId))
	s.cache.Delete(fmt.Sprintf("active-user-db-connection-%v", userId))
}
