package handler

import (
	"assignment-3/internal/db"
	"assignment-3/internal/form"
	"assignment-3/internal/render"
	"log"
	"net/http"
	"strconv"
)

func Venue(w http.ResponseWriter, r *http.Request) {

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
	if err := r.ParseForm(); err != nil {
		log.Println("Register:", err)
		return
	}

	if r.Method == http.MethodPost {
		form := form.New(r.PostForm)
		newVenue := db.Venue{
			Name: r.FormValue("name"),
			Type: r.FormValue("type"),
		}
		form.Required("name", "capacity", "type")

		capacity, err := strconv.Atoi(r.FormValue("capacity"))
		if err != nil {
			form.Errors.Add("capacity", "capacity accepts integers only")
		}
		if capacity < 1 {
			form.Errors.Add("capacity", "capacity must be more than 0")
		}

		if !form.Valid() {
			data := make(map[string]interface{})
			data["venue"] = newVenue
			data["user"] = user

			if err := render.Template(w, r, "venue.page.html", &render.TemplateData{
				Data: data,
				Form: form,
			}); err != nil {
				log.Println("Admin: Venue: ", err)
			}
			return
		}
		newVenue.Capacity = capacity
		db.Add(&newVenue)

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	if err := render.Template(w, r, "venue.page.html",
		&render.TemplateData{
			Data: make(map[string]interface{}),
			Form: form.New(nil)}); err != nil {
		log.Println("Admin: Venue: ", err)
		return
	}
}
