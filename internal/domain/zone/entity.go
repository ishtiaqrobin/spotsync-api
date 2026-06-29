package zone

import (
	"gorm.io/gorm"

	"github.com/ishtiaqrobin/spotsync-api/internal/domain/zone/dto"
)

// ParkingZone represents the parking_zones table in the database
type ParkingZone struct {
	gorm.Model
	Name          string  `json:"name" gorm:"type:varchar(100);not null"`
	Type          string  `json:"type" gorm:"type:varchar(20);not null"`
	TotalCapacity int     `json:"total_capacity" gorm:"not null"`
	PricePerHour  float64 `json:"price_per_hour" gorm:"not null"`
}

// TableName explicitly sets the table name
func (ParkingZone) TableName() string {
	return "parking_zones"
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
