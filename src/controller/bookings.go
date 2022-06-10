package controller

import (
	"class-booker/src/model"
	"class-booker/src/repository"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type BookingController struct {
	Repo *repository.Repository
}

func NewBookingController(repo *repository.Repository) *BookingController {
	return &BookingController{
		Repo: repo,
	}
}

func (c *BookingController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking model.Booking
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()

	isValid, err := isValidBooking(booking, c.Repo.Classes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if isValid {
		c.Repo.Bookings = append(c.Repo.Bookings, booking)
		response, _ := json.Marshal(booking)
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
		return
	}
}

func (c *BookingController) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	var booking model.Booking
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()

	for i, existingBooking := range c.Repo.Bookings {
		if strings.EqualFold(existingBooking.Name, booking.Name) && strings.EqualFold(existingBooking.ClassName, booking.ClassName) {
			isValid, err := isValidBooking(booking, c.Repo.Classes)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			if isValid {
				c.Repo.Bookings[i] = booking

				response, _ := json.Marshal(booking)
				w.WriteHeader(http.StatusOK)
				w.Write(response)
				return
			}
		}
	}

	http.Error(w, "booking for class "+booking.ClassName+" is not found for member "+booking.Name, http.StatusNotFound)
}

func (c *BookingController) GetBookings(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(c.Repo.Bookings)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *BookingController) GetBooking(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name := pathParams["member"]
	queryParams := r.URL.Query()
	className := queryParams.Get("class")

	for _, booking := range c.Repo.Bookings {
		if strings.EqualFold(booking.Name, name) && strings.EqualFold(booking.ClassName, className) {
			response, _ := json.Marshal(booking)
			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}
	}

	http.Error(w, "booking for class "+className+" is not found for member "+name, http.StatusNotFound)

}

func (c *BookingController) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name := pathParams["member"]
	queryParams := r.URL.Query()
	className := queryParams.Get("class")

	for i, booking := range c.Repo.Bookings {
		if strings.EqualFold(booking.Name, name) && strings.EqualFold(booking.ClassName, className) {
			c.Repo.Bookings = append(c.Repo.Bookings[:i], c.Repo.Bookings[i+1:]...)

			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "booking for class "+className+" is not found for member "+name, http.StatusNotFound)
}

func isValidBooking(booking model.Booking, classes []model.Class) (bool, error) {
	for _, class := range classes {
		if strings.EqualFold(class.Name, booking.ClassName) {
			startDate, _ := time.Parse("2006-01-02", class.StartDate)
			endDate, _ := time.Parse("2006-01-02", class.EndDate)
			bookingDate, _ := time.Parse("2006-01-02", booking.Date)

			if bookingDate.After(startDate) && bookingDate.Before(endDate) {
				return true, nil
			} else {
				return false, errors.New("class is not available on " + booking.Date)
			}
		}
	}

	return false, errors.New(booking.ClassName + " class does not exist")
}
