package repository

import "github.com/sangolariel/bookings/internal/models"

type DatabaseRepo interface {
	AddUsers() bool

	InsertReservation(res models.Reservation) error
}
