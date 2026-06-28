package models

import "time"

type Reservation struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	UserID       uint        `gorm:"not null" json:"user_id"`
	ZoneID       uint        `gorm:"not null" json:"zone_id"`
	LicensePlate string      `gorm:"size:15;not null" json:"license_plate"`
	Status       string      `gorm:"type:varchar(15);default:'active';check:status IN ('active','completed','cancelled')" json:"status"`
	User         User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Zone         ParkingZone `gorm:"foreignKey:ZoneID" json:"zone,omitempty"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
