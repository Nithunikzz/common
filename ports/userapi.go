package ports

import (
	//"github.com/google/uuid"
	"github.com/Nithunikzz/common/domain"
)

type UserAPIPort interface {
	CreateUser(user *domain.User) (*domain.UserIDResponse, error)
	DeleteUser(userid string) error
	GetUser(userid string) (*domain.User, error)
	UpdateUser(user *domain.User) (*domain.UserIDResponse, error)
	GetUserByOryID(oryid string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
