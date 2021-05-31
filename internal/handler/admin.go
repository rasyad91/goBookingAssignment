package handler

import (
	"assignment-3/internal/db"
	"assignment-3/internal/render"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	// get user in session
	user := getUser(r)
	// not logged in, cannot see bookings
	if user == nil || !user.IsAdmin {
		if err := render.Template(w, r, "restricted.page.html", &render.TemplateData{
			Data: data,
		}); err != nil {
			log.Println("Admin: Error parsing template: ", err)
		}
		return
	}
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Println("Admin: ", err)
		}
		venueID, err := strconv.Atoi(r.PostFormValue("venueID"))
		if err != nil {
			log.Println("Admin: Parsing String to int: ", err)
		}

		var wg sync.WaitGroup
		wg.Add(2)

		go func(id int) {
			if err := recover(); err != nil {
				fmt.Println("Admin: ", err)
			}
			db.DeleteBookingsFromBookingList(id)
			wg.Done()
		}(venueID)

		go func(id int) {
			if err := recover(); err != nil {
				fmt.Println("Admin: ", err)
			}
			db.DeleteBookingsFromUserBooking(id)
			wg.Done()
		}(venueID)

		if err := db.DeleteVenue(venueID); err != nil {
			log.Println("Admin: Error deleting venue: ", err)
		}
		wg.Wait()
	}

	data["user"] = user
	data["venue"] = db.VenueList

	if err := render.Template(w, r, "admin.page.html", &render.TemplateData{
		Data: data,
	}); err != nil {
		log.Println("Admin: Error parsing template: ", err)
	}

}
