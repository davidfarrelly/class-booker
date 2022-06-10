package controller

import (
	"class-booker/src/model"
	"class-booker/src/repository"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type ClassController struct {
	Repo *repository.Repository
}

func NewClassController(repo *repository.Repository) *ClassController {
	return &ClassController{
		Repo: repo,
	}
}

func (c *ClassController) UpdateClass(w http.ResponseWriter, r *http.Request) {
	var class model.Class
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&class); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()

	for i, existingClass := range c.Repo.Classes {
		if strings.EqualFold(existingClass.Name, class.Name) {
			c.Repo.Classes[i] = class

			response, _ := json.Marshal(class)
			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}
	}

	http.Error(w, class.Name+" class is not found", http.StatusNotFound)
}

func (c *ClassController) CreateClass(w http.ResponseWriter, r *http.Request) {
	var class model.Class
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&class); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()

	c.Repo.Classes = append(c.Repo.Classes, class)

	response, _ := json.Marshal(class)
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (c *ClassController) GetClasses(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(c.Repo.Classes)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *ClassController) GetClass(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["class"]

	for _, class := range c.Repo.Classes {
		if strings.EqualFold(class.Name, name) {
			response, _ := json.Marshal(class)
			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}
	}

	http.Error(w, name+" class is not found", http.StatusNotFound)

}

func (c *ClassController) DeleteClass(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["class"]

	for i, class := range c.Repo.Classes {
		if strings.EqualFold(class.Name, name) {
			c.Repo.Classes = append(c.Repo.Classes[:i], c.Repo.Classes[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, name+" class is not found", http.StatusNotFound)
}
