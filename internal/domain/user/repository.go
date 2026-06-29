package user

import (
	"errors"

	"gorm.io/gorm"
)

// ErrEmailAlreadyExists is returned when a user with the given email already exists
var ErrEmailAlreadyExists = errors.New("user with this email already exists")

// Repository defines the interface for user data access
type Repository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id uint) (*User, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new user Repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Create inserts a new user into the database
func (r *repository) Create(user *User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrEmailAlreadyExists
		}
		return result.Error
	}
	return nil
}

// FindByEmail finds a user by their email address
func (r *repository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByID finds a user by their ID
func (r *repository) FindByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
