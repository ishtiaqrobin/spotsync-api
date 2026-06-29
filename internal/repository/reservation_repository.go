package repository

import (
	"errors"

	"github.com/ishtiaqrobin/spotsync-api/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrZoneFull = errors.New("parking zone is full")
var ErrZoneNotFound = errors.New("parking zone not found")

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

// CreateWithLock safely creates a reservation by locking the zone row first,
// preventing the "EV Spot Bottleneck" race condition described in the README.
func (r *ReservationRepository) CreateWithLock(reservation *models.Reservation) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var zone models.ParkingZone

		// 1. Lock the parking zone row (SELECT ... FOR UPDATE)
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&zone, reservation.ZoneID).Error; err != nil {
			return ErrZoneNotFound
		}

		// 2. Count current active reservations for this zone (inside the same transaction)
		var activeCount int64
		if err := tx.Model(&models.Reservation{}).
			Where("zone_id = ? AND status = ?", zone.ID, "active").
			Count(&activeCount).Error; err != nil {
			return err
		}

		// 3. Check capacity
		if activeCount >= int64(zone.TotalCapacity) {
			return ErrZoneFull
		}

		// 4. Create the reservation (status defaults to "active" via GORM model tag)
		if err := tx.Create(reservation).Error; err != nil {
			return err
		}

		return nil // commits transaction
	})
}

func (r *ReservationRepository) FindByUserID(userID uint) ([]models.Reservation, error) {
	var reservations []models.Reservation
	err := r.db.Preload("Zone").
		Where("user_id = ?", userID).
		Find(&reservations).Error
	return reservations, err
}

func (r *ReservationRepository) FindAll() ([]models.Reservation, error) {
	var reservations []models.Reservation
	err := r.db.Preload("User").Preload("Zone").Find(&reservations).Error
	return reservations, err
}

func (r *ReservationRepository) FindByID(id uint) (*models.Reservation, error) {
	var reservation models.Reservation
	err := r.db.First(&reservation, id).Error
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *ReservationRepository) UpdateStatus(reservation *models.Reservation, status string) error {
	return r.db.Model(reservation).Update("status", status).Error
}
