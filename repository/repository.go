package repository

import (
	"errors"

	"gorm.io/gorm"
)

var (
	errInvalidUser = errors.New("invalid email or password")
)

type Repo struct {
	db *gorm.DB
}

type RepoInterface interface {
	UserRepo
	PhotoRepo
	CommentRepo
	SocialMediaRepo
}

// constructor function
func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db} // handle dependencies
}
