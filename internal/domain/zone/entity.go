package zone

import (
	"gorm.io/gorm"

	"github.com/ishtiaqrobin/spotsync-api/internal/domain/zone/dto"
)

// ParkingZone represents the parking_zones table in the database
type ParkingZone struct {
	gorm.Model
	Name          string  `json:"name" gorm:"type:varchar(100);not null"`
	Type          string  `json:"type" gorm:"type:varchar(20);not null;check:type IN ('general','ev_charging','covered')"`
	TotalCapacity int     `json:"total_capacity" gorm:"not null;check:total_capacity > 0"`
	PricePerHour  float64 `json:"price_per_hour" gorm:"not null;check:price_per_hour > 0"`

	// Relations
	Reservations []Reservation `json:"reservations,omitempty" gorm:"foreignKey:ZoneID"`
}

// Reservation is a minimal reference for the relation (full model is in reservation domain)
type Reservation struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	ZoneID       uint   `json:"zone_id"`
	LicensePlate string `json:"license_plate"`
	Status       string `json:"status"`
}

// ToResponse converts a ParkingZone entity to a DTO response
func (z *ParkingZone) ToResponse(availableSpots int) *dto.Response {
	return &dto.Response{
		ID:             z.ID,
		Name:           z.Name,
		Type:           z.Type,
		TotalCapacity:  z.TotalCapacity,
		AvailableSpots: availableSpots,
		PricePerHour:   z.PricePerHour,
		CreatedAt:      z.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:      z.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
