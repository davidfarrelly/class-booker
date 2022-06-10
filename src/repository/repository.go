package repository

import "class-booker/src/model"

type Repository struct {
	Classes  []model.Class
	Bookings []model.Booking
}

func InitNewRepository() *Repository {
	return &Repository{
		Classes:  []model.Class{},
		Bookings: []model.Booking{},
	}
}
