package concurrency

import (
	"library_management/services"
	"time"
)

func StartReservationWorker(lib *services.Library, ch chan services.ReservationRequest) {
	for req := range ch {
		go func(req services.ReservationRequest) {
			time.Sleep(5 * time.Second)
			lib.Mu.Lock()
			defer lib.Mu.Unlock()
			id := req.BookID
			book := lib.Books[id]
			if book.Status == "reserved" {
				book.Status = "available"
				lib.Books[id] = book
			}
		}(req)
	}
}
