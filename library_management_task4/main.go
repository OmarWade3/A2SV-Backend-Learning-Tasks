package main

import (
	"library_management/concurrency"
	"library_management/controllers"
	"library_management/services"
)

func main() {
	reservationChan := make(chan services.ReservationRequest)

	library := services.NewLibrary(reservationChan)

	go concurrency.StartReservationWorker(library, reservationChan)

	controllers.RunLibrary(library)
}
