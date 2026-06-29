package reservation

import (
	"errors"

	"gorm.io/gorm"
)

// ErrZoneFull is returned when a parking zone has no available spots
var ErrZoneFull = errors.New("parking zone is full")

// ErrZoneNotFound is returned when the referenced parking zone does not exist
var ErrZoneNotFound = errors.New("parking zone not found")

// ErrReservationNotFound is returned when a reservation is not found
var ErrReservationNotFound = errors.New("reservation not found")

// ErrForbidden is returned when a user tries to access another user's reservation
var ErrForbidden = errors.New("you are not allowed to perform this action")

// Repository defines the interface for reservation data access
type Repository interface {
	CreateReservation(reservation *Reservation) error
	FindByUserID(userID uint) ([]Reservation, error)
	FindAll() ([]Reservation, error)
	FindByID(id uint) (*Reservation, error)
	UpdateStatus(reservation *Reservation, status string) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new reservation Repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// CreateReservation safely creates a reservation with capacity check
func (r *repository) CreateReservation(reservation *Reservation) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var zone Zone

		// 1. Find the parking zone
		if err := tx.First(&zone, reservation.ZoneID).Error; err != nil {
			return ErrZoneNotFound
		}

		// 2. Count current active reservations for this zone
		var activeCount int64
		if err := tx.Model(&Reservation{}).
			Where("zone_id = ? AND status = ?", zone.ID, "active").
			Count(&activeCount).Error; err != nil {
			return err
		}

		// 3. Check capacity
		if activeCount >= int64(zone.TotalCapacity) {
			return ErrZoneFull
		}

		// 4. Create the reservation
		if err := tx.Create(reservation).Error; err != nil {
			return err
		}

		return nil // commits transaction
	})
}

// FindByUserID retrieves all reservations for a given user, with zone info preloaded
func (r *repository) FindByUserID(userID uint) ([]Reservation, error) {
	var reservations []Reservation
	err := r.db.Preload("Zone").
		Where("user_id = ?", userID).
		Find(&reservations).Error
	return reservations, err
}

// FindAll retrieves all reservations with user and zone info preloaded
func (r *repository) FindAll() ([]Reservation, error) {
	var reservations []Reservation
	err := r.db.Preload("User").Preload("Zone").Find(&reservations).Error
	return reservations, err
}

// FindByID retrieves a reservation by its ID
func (r *repository) FindByID(id uint) (*Reservation, error) {
	var reservation Reservation
	err := r.db.First(&reservation, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrReservationNotFound
		}
		return nil, err
	}
	return &reservation, nil
}

// UpdateStatus updates the status of a reservation
func (r *repository) UpdateStatus(reservation *Reservation, status string) error {
	return r.db.Model(reservation).Update("status", status).Error
}
