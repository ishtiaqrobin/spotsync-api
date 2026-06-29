package zone

import (
	"errors"

	"github.com/ishtiaqrobin/spotsync-api/internal/domain/zone/dto"
)

// ErrZoneNotFound is returned when a parking zone is not found
var ErrZoneNotFound = errors.New("parking zone not found")

type service struct {
	repo Repository
}

// NewService creates a new parking zone service
func NewService(repo Repository) *service {
	return &service{repo: repo}
}

// CreateZone creates a new parking zone
func (s *service) CreateZone(req dto.CreateZoneRequest) (*dto.Response, error) {
	zone := ParkingZone{
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	if err := s.repo.Create(&zone); err != nil {
		return nil, err
	}

	// New zone has no reservations yet, so available = total capacity
	return zone.ToResponse(zone.TotalCapacity), nil
}

// GetAllZones retrieves all zones with dynamically calculated available spots
func (s *service) GetAllZones() ([]dto.Response, error) {
	zones, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []dto.Response
	for _, z := range zones {
		activeCount, err := s.repo.CountActiveReservations(z.ID)
		if err != nil {
			return nil, err
		}

		available := z.TotalCapacity - int(activeCount)
		result = append(result, *z.ToResponse(available))
	}

	return result, nil
}

// GetZoneByID retrieves a single zone with calculated available spots
func (s *service) GetZoneByID(id uint) (*dto.Response, error) {
	zone, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	activeCount, err := s.repo.CountActiveReservations(zone.ID)
	if err != nil {
		return nil, err
	}

	available := zone.TotalCapacity - int(activeCount)
	return zone.ToResponse(available), nil
}
