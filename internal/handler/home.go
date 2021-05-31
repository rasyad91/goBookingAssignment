package handler

import (
	"assignment-3/internal/db"
	"assignment-3/internal/form"
	"assignment-3/internal/render"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["venues"] = db.VenueList

	user := getUser(r)
	if user != nil {
		data["user"] = user

		if r.Method == http.MethodPost {
			r.ParseForm()
			venueID, err := strconv.Atoi(r.PostFormValue("venueID"))
			if err != nil {
				log.Println("Home:", err)
			}
			date := r.PostFormValue(fmt.Sprintf("date%d", venueID))
			if !db.BookingIsAvailable(venueID, date) {
				form := form.New(r.Form)
				form.Errors.Add("date", fmt.Sprintf("Date selected for \"%s\" has already been booked! Please select another date", db.GetByID(venueID).Name))
				if err := render.Template(w, r, "home.page.html", &render.TemplateData{Data: data, Form: form}); err != nil {
					log.Println("Home: ", err)
				}
				return
			}

			newBooking := db.Booking{
				CustomerID: user.CustomerID,
				VenueID:    venueID,
				Date:       date,
			}
			bookingID := db.NewBooking(newBooking)
			user.BookingsID = append(user.BookingsID, bookingID)
			form := form.New(r.Form)
			form.Errors.Add("success", fmt.Sprintf("Booking for \"%s\" on \"%s\" successful!", db.GetByID(venueID).Name, date))

			if err := render.Template(w, r, "home.page.html", &render.TemplateData{Data: data, Form: form}); err != nil {
				log.Println("Home: ", err)
			}
			return

		}

		if err := render.Template(w, r, "home.page.html", &render.TemplateData{Data: data, Form: form.New(nil)}); err != nil {
			log.Println("Home: ", err)
		}

	} else {
		if err := render.Template(w, r, "home.page.html",
			&render.TemplateData{
				Data: data,
				Form: form.New(nil)}); err != nil {
			log.Println("Home: ", err)
		}
	}
}

func getUser(r *http.Request) (user *db.User) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil
	}

	if username, ok := db.Sessions[cookie.Value]; ok {
		user = db.Users[username]
	}
	return
}
