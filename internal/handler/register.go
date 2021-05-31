package handler

import (
	"assignment-3/internal/db"
	"assignment-3/internal/form"
	"assignment-3/internal/render"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	if getUser(r) != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println("Register:", err)
		return
	}

	if r.Method == http.MethodPost {
		newUser := db.User{
			Firstname: r.FormValue("firstname"),
			Lastname:  r.FormValue("lastname"),
			Username:  r.FormValue("username"),
			Password:  []byte(r.FormValue("password")),
		}
		form := form.New(r.PostForm)

		form.Required("firstname", "lastname", "username", "password")
		if form.ExistingUser() {
			form.Errors.Add("username", "Username already in use")
		}
		if !form.Valid() {
			data := make(map[string]interface{})
			data["register"] = newUser

			if err := render.Template(w, r, "register.page.html", &render.TemplateData{
				Data: data,
				Form: form,
			}); err != nil {
				log.Println("Registration: ", err)
			}
			return
		}

		if err := db.CreateNewUser(&newUser); err != nil {
			log.Println("Registration: ", err)
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if err := render.Template(w, r, "register.page.html",
		&render.TemplateData{
			Data: make(map[string]interface{}),
			Form: form.New(nil)}); err != nil {
		log.Println("Registration: ", err)
		return
	}
}
