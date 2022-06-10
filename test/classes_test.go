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

func TestCreateClass(t *testing.T) {
	var body = []byte(`{"name":"yoga","capacity":"10","startDate":"2022-06-10","endDate":"2022-07-10"}`)
	req, _ := http.NewRequest("POST", "/classes", bytes.NewBuffer(body))
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	controller := controller.NewClassController(repo)

	handler := http.HandlerFunc(controller.CreateClass)
	handler.ServeHTTP(r, req)

	assert.Equal(t, http.StatusCreated, r.Code)
	assert.Equal(t, 1, len(controller.Repo.Classes))
}

func TestUpdateClass(t *testing.T) {
	var body = []byte(`{"name":"yoga","capacity":"20","startDate":"2022-06-10","endDate":"2022-07-10"}`)
	req, _ := http.NewRequest("PUT", "/classes", bytes.NewBuffer(body))
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Classes = append(repo.Classes, model.Class{Name: "yoga", Capacity: "10", StartDate: "2022-06-10", EndDate: "2022-07-10"})

	controller := controller.NewClassController(repo)

	handler := http.HandlerFunc(controller.UpdateClass)
	handler.ServeHTTP(r, req)

	assert.Equal(t, http.StatusOK, r.Code)
	assert.Equal(t, "20", controller.Repo.Classes[0].Capacity)
}

func TestUpdateClassNotExist(t *testing.T) {
	var body = []byte(`{"name":"crossfit","capacity":"20","startDate":"2022-06-10","endDate":"2022-07-10"}`)
	req, _ := http.NewRequest("PUT", "/classes", bytes.NewBuffer(body))
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Classes = append(repo.Classes, model.Class{Name: "yoga", Capacity: "10", StartDate: "2022-06-10", EndDate: "2022-07-10"})

	controller := controller.NewClassController(repo)

	handler := http.HandlerFunc(controller.UpdateClass)
	handler.ServeHTTP(r, req)

	var expectedError = "crossfit class is not found"
	assert.Equal(t, http.StatusNotFound, r.Code)
	assert.Contains(t, r.Body.String(), expectedError)
}

func TestGetClasses(t *testing.T) {
	var expectedResponse = []byte(`[{"name":"yoga","capacity":"10","startDate":"2022-06-10","endDate":"2022-07-10"}]`)
	req, _ := http.NewRequest("GET", "/classes", nil)
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Classes = append(repo.Classes, model.Class{Name: "yoga", Capacity: "10", StartDate: "2022-06-10", EndDate: "2022-07-10"})

	controller := controller.NewClassController(repo)

	handler := http.HandlerFunc(controller.GetClasses)
	handler.ServeHTTP(r, req)

	assert.Equal(t, http.StatusOK, r.Code)
	assert.Equal(t, expectedResponse, r.Body.Bytes())
}

func TestGetClass(t *testing.T) {
	var expectedResponse = []byte(`{"name":"yoga","capacity":"10","startDate":"2022-06-10","endDate":"2022-07-10"}`)
	req, _ := http.NewRequest("GET", "/classes/class", nil)
	req = mux.SetURLVars(req, map[string]string{"class": "yoga"})
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Classes = append(repo.Classes, model.Class{Name: "yoga", Capacity: "10", StartDate: "2022-06-10", EndDate: "2022-07-10"})

	controller := controller.NewClassController(repo)

	handler := http.HandlerFunc(controller.GetClass)
	handler.ServeHTTP(r, req)

	assert.Equal(t, http.StatusOK, r.Code)
	assert.Equal(t, expectedResponse, r.Body.Bytes())
}

func TestGetClassNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/classes/class", nil)
	req = mux.SetURLVars(req, map[string]string{"class": "crossfit"})
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Classes = append(repo.Classes, model.Class{Name: "yoga", Capacity: "10", StartDate: "2022-06-10", EndDate: "2022-07-10"})

	controller := controller.NewClassController(repo)

	handler := http.HandlerFunc(controller.GetClass)
	handler.ServeHTTP(r, req)

	var expectedError = "crossfit class is not found"
	assert.Equal(t, http.StatusNotFound, r.Code)
	assert.Contains(t, r.Body.String(), expectedError)
}

func TestDeleteClass(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/classes/class", nil)
	req = mux.SetURLVars(req, map[string]string{"class": "yoga"})
	r := httptest.NewRecorder()

	repo := repository.InitNewRepository()
	repo.Classes = append(repo.Classes, model.Class{Name: "yoga", Capacity: "10", StartDate: "2022-06-10", EndDate: "2022-07-10"})

	controller := controller.NewClassController(repo)

	handler := http.HandlerFunc(controller.DeleteClass)
	handler.ServeHTTP(r, req)

	assert.Equal(t, http.StatusOK, r.Code)
	assert.Equal(t, 0, len(controller.Repo.Classes))
}
