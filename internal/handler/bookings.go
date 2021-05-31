package handler

import (
	"assignment-3/internal/db"
	"assignment-3/internal/form"
	"assignment-3/internal/render"
	"log"
	"net/http"
	"strconv"
)

func Bookings(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	// get user in session
	user := getUser(r)
	// not logged in, cannot see bookings
	if user == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	data["user"] = user

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Println("Bookings: ", err)
		}
		bookingID, err := strconv.Atoi(r.PostFormValue("bookingID"))
		if err != nil {
			log.Println("Booking: Parsing String to int: ", err)
		}
		if err := user.CancelBooking(bookingID); err != nil {
			log.Fatalln(err)
		}

		delete(db.BookingList, bookingID)
	}
	usersBookings := []db.Booking{}

	for _, v := range user.BookingsID {
		usersBookings = append(usersBookings, db.BookingList[v])
	}
	data["bookinglist"] = usersBookings

	if err := render.Template(w, r, "bookings.page.html", &render.TemplateData{Data: data, Form: form.New(nil)}); err != nil {
		log.Println("Bookings: ", err)
	}

}
