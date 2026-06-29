package reservation

import (
	"github.com/ishtiaqrobin/spotsync-api/internal/domain/reservation/dto"
)

type service struct {
	repo Repository
}

// NewService creates a new reservation service
func NewService(repo Repository) *service {
	return &service{repo: repo}
}

// CreateReservation creates a new reservation with zone capacity check
func (s *service) CreateReservation(userID uint, req dto.CreateReservationRequest) (*dto.Response, error) {
	reservation := Reservation{
		UserID:       userID,
		ZoneID:       req.ZoneID,
		LicensePlate: req.LicensePlate,
		Status:       "active",
	}

	if err := s.repo.CreateReservation(&reservation); err != nil {
		return nil, err
	}

	return reservation.ToResponse(), nil
}

// GetMyReservations retrieves all reservations for a specific user
func (s *service) GetMyReservations(userID uint) ([]dto.MyReservationResponse, error) {
	reservations, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var result []dto.MyReservationResponse
	for _, r := range reservations {
		result = append(result, *r.ToMyResponse())
	}

	return result, nil
}

// GetAllReservations retrieves all reservations in the system (admin only)
func (s *service) GetAllReservations() ([]dto.Response, error) {
	reservations, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []dto.Response
	for _, r := range reservations {
		result = append(result, *r.ToResponse())
	}

	return result, nil
}

// CancelReservation cancels a reservation (only owner or admin can cancel)
func (s *service) CancelReservation(reservationID uint, userID uint, userRole string) error {
	reservation, err := s.repo.FindByID(reservationID)
	if err != nil {
		return err
	}

	// Driver can only cancel their own reservation
	if userRole != "admin" && reservation.UserID != userID {
		return ErrForbidden
	}

	return s.repo.UpdateStatus(reservation, "cancelled")
}
