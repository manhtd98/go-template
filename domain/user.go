package domain

import (
	"context"

	"gorm.io/gorm"
)

const (
	CollectionUser = "users"
)

type User struct {
	gorm.Model
	ID   string
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}
