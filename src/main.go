package main

import (
	"class-booker/src/controller"
	"class-booker/src/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Initializing Repository")
	repo := repository.InitNewRepository()
	classController := controller.NewClassController(repo)
	bookingController := controller.NewBookingController(repo)

	router := mux.NewRouter()

	fmt.Println("Adding Routes")
	addRoutes(router, *classController, *bookingController)

	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func addRoutes(router *mux.Router, classController controller.ClassController, bookingController controller.BookingController) {
	router.HandleFunc("/classes", classController.CreateClass).Methods("POST")
	router.HandleFunc("/classes", classController.UpdateClass).Methods("PUT")
	router.HandleFunc("/classes/{class}", classController.GetClass).Methods("GET")
	router.HandleFunc("/classes", classController.GetClasses).Methods("GET")
	router.HandleFunc("/classes/{class}", classController.DeleteClass).Methods("DELETE")

	router.HandleFunc("/bookings", bookingController.CreateBooking).Methods("POST")
	router.HandleFunc("/bookings", bookingController.UpdateBooking).Methods("PUT")
	router.HandleFunc("/bookings/{member}", bookingController.GetBooking).Methods("GET")
	router.HandleFunc("/bookings", bookingController.GetBookings).Methods("GET")
	router.HandleFunc("/bookings/{member}", bookingController.DeleteBooking).Methods("DELETE")
}
