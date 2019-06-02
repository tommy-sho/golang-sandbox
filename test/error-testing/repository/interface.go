package repository

import (
	"time"

	"golang.org/x/xerrors"

	uuid "github.com/satori/go.uuid"
	"github.com/tommy-sho/golang-sandbox/test/error-testing/domain"
	validator2 "gopkg.in/go-playground/validator.v9"
)

var validator = validator2.New()

var (
	ValidationErr = xerrors.New("invalid parameter")
	NotExistErr   = xerrors.New("user does't exits")
)

type Repository interface {
	CreateUser(user *domain.User) error
	GetUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
}

func NewRepository() Repository {
	users := make(map[string]*domain.User)
	return repositoryImpl{
		users: users,
	}
}

type repositoryImpl struct {
	users map[string]*domain.User
}

func (r repositoryImpl) CreateUser(user *domain.User) error {
	err := validator.Struct(*user)
	if err != nil {
		return xerrors.Errorf("CreateUser error: %w", ValidationErr)
	}

	id := uuid.NewV4().String()
	now := time.Now()
	user.ID, user.CreatedAt, user.UpdatedAt = id, now, now
	r.users[user.ID] = user
	return nil
}

func (r repositoryImpl) GetUsers() ([]*domain.User, error) {
	users := make([]*domain.User, len(r.users))
	count := 0
	for _, v := range r.users {
		users[count] = v
		count += 1
	}

	return users, nil
}

func (r repositoryImpl) UpdateUser(user *domain.User) error {
	u := r.users[user.ID]
	if u.ID == "" {
		return xerrors.Errorf("UpdateUser error: %w", NotExistErr)
	}

	return nil
}

func (r repositoryImpl) DeleteUser(id string) error {
	delete(r.users, id)
	return nil
}
