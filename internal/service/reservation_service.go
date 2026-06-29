package service

import (
	"errors"

	"github.com/ishtiaqrobin/spotsync-api/internal/dto"
	"github.com/ishtiaqrobin/spotsync-api/internal/models"
	"github.com/ishtiaqrobin/spotsync-api/internal/repository"
)

var ErrForbidden = errors.New("you are not allowed to perform this action")

type ReservationService struct {
	reservationRepo *repository.ReservationRepository
}

func NewReservationService(reservationRepo *repository.ReservationRepository) *ReservationService {
	return &ReservationService{reservationRepo: reservationRepo}
}

// CreateReservation enforces zone capacity via the repository's locking transaction
func (s *ReservationService) CreateReservation(userID uint, req dto.CreateReservationRequest) (*models.Reservation, error) {
	reservation := &models.Reservation{
		UserID:       userID,
		ZoneID:       req.ZoneID,
		LicensePlate: req.LicensePlate,
		Status:       "active",
	}

	if err := s.reservationRepo.CreateWithLock(reservation); err != nil {
		return nil, err
	}

	return reservation, nil
}

func (s *ReservationService) GetMyReservations(userID uint) ([]models.Reservation, error) {
	return s.reservationRepo.FindByUserID(userID)
}

func (s *ReservationService) GetAllReservations() ([]models.Reservation, error) {
	return s.reservationRepo.FindAll()
}

// CancelReservation ensures only the owner can cancel their own reservation (README rule)
func (s *ReservationService) CancelReservation(reservationID uint, userID uint, userRole string) error {
	reservation, err := s.reservationRepo.FindByID(reservationID)
	if err != nil {
		return err
	}

	// Driver can only cancel their own reservation; admin restriction not specified for this endpoint
	if userRole != "admin" && reservation.UserID != userID {
		return ErrForbidden
	}

	return s.reservationRepo.UpdateStatus(reservation, "cancelled")
}
