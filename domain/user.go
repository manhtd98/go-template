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
	ID   int32           `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name     string             `json:"name"`
	UserName string			 `json:"username" gorm:"index:id_username"`
	Email    string             `json:"email" gorm:"index:idx_name;unique;not null;type:varchar(100);default:null"`
	Password string             `json:"password"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}
