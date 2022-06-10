package test

import (
	"bytes"
	"class-booker/src/controller"
	"class-booker/src/model"
	"class-booker/src/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	var body = []byte(`{"name":"david","className":"yoga","date":"2022-06-20"}`)
	req, _ := http.NewRequest("POST", "/bookings", bytes.NewBuffer(body))
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Classes = append(repo.Classes, model.Class{Name: "yoga", Capacity: "10", StartDate: "2022-06-10", EndDate: "2022-07-10"})
	controller := controller.NewBookingController(repo)

	handler := http.HandlerFunc(controller.CreateBooking)
	handler.ServeHTTP(r, req)

	assert.Equal(t, http.StatusCreated, r.Code)
	assert.Equal(t, 1, len(controller.Repo.Bookings))
}

func TestCreateBookingInvalidDate(t *testing.T) {
	var body = []byte(`{"name":"david","className":"yoga","date":"2022-07-11"}`)
	req, _ := http.NewRequest("POST", "/bookings", bytes.NewBuffer(body))
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Classes = append(repo.Classes, model.Class{Name: "yoga", Capacity: "10", StartDate: "2022-06-10", EndDate: "2022-07-10"})
	controller := controller.NewBookingController(repo)

	handler := http.HandlerFunc(controller.CreateBooking)
	handler.ServeHTTP(r, req)

	var expectedError = "class is not available on 2022-07-11"
	assert.Equal(t, http.StatusBadRequest, r.Code)
	assert.Contains(t, r.Body.String(), expectedError)
}

func TestCreateBookingClassNotExist(t *testing.T) {
	var body = []byte(`{"name":"david","className":"crossfit","date":"2022-07-11"}`)
	req, _ := http.NewRequest("POST", "/bookings", bytes.NewBuffer(body))
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Classes = append(repo.Classes, model.Class{Name: "yoga", Capacity: "10", StartDate: "2022-06-10", EndDate: "2022-07-10"})
	controller := controller.NewBookingController(repo)

	handler := http.HandlerFunc(controller.CreateBooking)
	handler.ServeHTTP(r, req)

	var expectedError = "crossfit class does not exist"
	assert.Equal(t, http.StatusBadRequest, r.Code)
	assert.Contains(t, r.Body.String(), expectedError)
}

func TestUpdateBooking(t *testing.T) {
	var body = []byte(`{"name":"david","className":"yoga","date":"2022-06-22"}`)
	req, _ := http.NewRequest("PUT", "/bookings", bytes.NewBuffer(body))
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Classes = append(repo.Classes, model.Class{Name: "yoga", Capacity: "10", StartDate: "2022-06-10", EndDate: "2022-07-10"})
	repo.Bookings = append(repo.Bookings, model.Booking{Name: "david", ClassName: "yoga", Date: "2022-06-20"})
	controller := controller.NewBookingController(repo)

	handler := http.HandlerFunc(controller.UpdateBooking)
	handler.ServeHTTP(r, req)

	assert.Equal(t, http.StatusOK, r.Code)
	assert.Equal(t, "2022-06-22", controller.Repo.Bookings[0].Date)
}

func TestGetBookings(t *testing.T) {
	var expectedResponse = []byte(`[{"name":"david","className":"yoga","date":"2022-06-20"}]`)
	req, _ := http.NewRequest("GET", "/bookings", nil)
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Bookings = append(repo.Bookings, model.Booking{Name: "david", ClassName: "yoga", Date: "2022-06-20"})

	controller := controller.NewBookingController(repo)

	handler := http.HandlerFunc(controller.GetBookings)
	handler.ServeHTTP(r, req)

	assert.Equal(t, http.StatusOK, r.Code)
	assert.Equal(t, expectedResponse, r.Body.Bytes())
}

func TestGetBooking(t *testing.T) {
	var expectedResponse = []byte(`{"name":"david","className":"yoga","date":"2022-06-20"}`)
	req, _ := http.NewRequest("GET", "/bookings/member?class=yoga", nil)
	req = mux.SetURLVars(req, map[string]string{"member": "david"})
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Bookings = append(repo.Bookings, model.Booking{Name: "david", ClassName: "yoga", Date: "2022-06-20"})

	controller := controller.NewBookingController(repo)

	handler := http.HandlerFunc(controller.GetBooking)
	handler.ServeHTTP(r, req)

	assert.Equal(t, http.StatusOK, r.Code)
	assert.Equal(t, expectedResponse, r.Body.Bytes())
}

func TestGetBookingNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/bookings/member?class=crossfit", nil)
	req = mux.SetURLVars(req, map[string]string{"member": "david"})
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Bookings = append(repo.Bookings, model.Booking{Name: "david", ClassName: "yoga", Date: "2022-06-20"})

	controller := controller.NewBookingController(repo)

	handler := http.HandlerFunc(controller.GetBooking)
	handler.ServeHTTP(r, req)

	var expectedError = "booking for class crossfit is not found for member david"
	assert.Equal(t, http.StatusNotFound, r.Code)
	assert.Contains(t, r.Body.String(), expectedError)
}

func TestDeleteBooking(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/bookings/member?class=yoga", nil)
	req = mux.SetURLVars(req, map[string]string{"member": "david"})
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Bookings = append(repo.Bookings, model.Booking{Name: "david", ClassName: "yoga", Date: "2022-06-20"})

	controller := controller.NewBookingController(repo)

	handler := http.HandlerFunc(controller.DeleteBooking)
	handler.ServeHTTP(r, req)

	assert.Equal(t, http.StatusOK, r.Code)
	assert.Equal(t, 0, len(controller.Repo.Classes))
}
