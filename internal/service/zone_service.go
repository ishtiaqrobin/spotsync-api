package service

import (
	"github.com/ishtiaqrobin/spotsync-api/internal/dto"
	"github.com/ishtiaqrobin/spotsync-api/internal/models"
	"github.com/ishtiaqrobin/spotsync-api/internal/repository"
)

type ZoneService struct {
	zoneRepo *repository.ZoneRepository
}

func NewZoneService(zoneRepo *repository.ZoneRepository) *ZoneService {
	return &ZoneService{zoneRepo: zoneRepo}
}

func (s *ZoneService) CreateZone(req dto.CreateZoneRequest) (*models.ParkingZone, error) {
	zone := &models.ParkingZone{
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	if err := s.zoneRepo.Create(zone); err != nil {
		return nil, err
	}

	return zone, nil
}

// GetAllZones returns all zones with dynamically calculated available_spots
func (s *ZoneService) GetAllZones() ([]dto.ZoneResponse, error) {
	zones, err := s.zoneRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []dto.ZoneResponse
	for _, zone := range zones {
		activeCount, err := s.zoneRepo.CountActiveReservations(zone.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, dto.ZoneResponse{
			ID:             zone.ID,
			Name:           zone.Name,
			Type:           zone.Type,
			TotalCapacity:  zone.TotalCapacity,
			AvailableSpots: zone.TotalCapacity - int(activeCount),
			PricePerHour:   zone.PricePerHour,
			CreatedAt:      zone.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return result, nil
}

// GetZoneByID returns a single zone with available_spots calculated
func (s *ZoneService) GetZoneByID(id uint) (*dto.ZoneResponse, error) {
	zone, err := s.zoneRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	activeCount, err := s.zoneRepo.CountActiveReservations(zone.ID)
	if err != nil {
		return nil, err
	}

	return &dto.ZoneResponse{
		ID:             zone.ID,
		Name:           zone.Name,
		Type:           zone.Type,
		TotalCapacity:  zone.TotalCapacity,
		AvailableSpots: zone.TotalCapacity - int(activeCount),
		PricePerHour:   zone.PricePerHour,
		CreatedAt:      zone.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
