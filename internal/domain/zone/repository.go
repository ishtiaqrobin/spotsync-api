package zone

import (
	"errors"

	"gorm.io/gorm"
)

// Repository defines the interface for parking zone data access
type Repository interface {
	Create(zone *ParkingZone) error
	FindAll() ([]ParkingZone, error)
	FindByID(id uint) (*ParkingZone, error)
	CountActiveReservations(zoneID uint) (int64, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new parking zone Repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Create inserts a new parking zone into the database
func (r *repository) Create(zone *ParkingZone) error {
	return r.db.Create(zone).Error
}

// FindAll retrieves all parking zones
func (r *repository) FindAll() ([]ParkingZone, error) {
	var zones []ParkingZone
	err := r.db.Find(&zones).Error
	return zones, err
}

// FindByID retrieves a parking zone by its ID
func (r *repository) FindByID(id uint) (*ParkingZone, error) {
	var zone ParkingZone
	err := r.db.First(&zone, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrZoneNotFound
		}
		return nil, err
	}
	return &zone, nil
}

// CountActiveReservations counts active reservations for a given zone
func (r *repository) CountActiveReservations(zoneID uint) (int64, error) {
	var count int64
	err := r.db.Model(&Reservation{}).
		Where("zone_id = ? AND status = ?", zoneID, "active").
		Count(&count).Error
	return count, err
}
