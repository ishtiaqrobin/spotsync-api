package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/ishtiaqrobin/spotsync-api/internal/domain/user/dto"
)

// User represents the users table in the database
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`
	Role     string `json:"role" gorm:"type:varchar(10);default:'driver'"`

	// Relations
	Reservations []Reservation `json:"reservations,omitempty" gorm:"foreignKey:UserID"`
}

// Reservation is a minimal reference for the relation (full model is in reservation domain)
type Reservation struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	ZoneID       uint   `json:"zone_id"`
	LicensePlate string `json:"license_plate"`
	Status       string `json:"status"`
}

// hashPassword hashes the user's password using bcrypt
func (u *User) hashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// checkPassword verifies a plaintext password against the stored hash
func (u *User) checkPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// ToResponse converts a User entity to a safe DTO response (no password)
func (u *User) ToResponse() *dto.Response {
	return &dto.Response{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
