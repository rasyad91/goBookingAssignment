package handler

import (
	"assignment-3/internal/db"
	"assignment-3/internal/form"
	"assignment-3/internal/render"
	"log"
	"net/http"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

func LogIn(w http.ResponseWriter, r *http.Request) {

	if getUser(r) != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("Login:", err)
		return
	}

	if r.Method == http.MethodPost {
		user := db.User{
			Username: r.PostFormValue("username"),
			Password: []byte(r.PostFormValue("password")),
		}

		form := form.New(r.PostForm)
		form.Required("username", "password")
		if !form.ExistingUser() {
			form.Errors.Add("username", "Username and/or password do not match")

		} else {
			if err := bcrypt.CompareHashAndPassword(db.Users[user.Username].Password, user.Password); err != nil {
				form.Errors.Add("username", "Username and/or password do not match")
			}
		}
		if !form.Valid() {

			data := make(map[string]interface{})
			data["login"] = user
			if err := render.Template(w, r, "/login.page.html", &render.TemplateData{
				Data: data,
				Form: form,
			}); err != nil {
				log.Println("Login: ", err)
			}
			return
		}

		id, err := uuid.NewRandom()
		if err != nil {
			log.Println("Login:", err)
		}
		cookie := &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		db.Sessions[cookie.Value] = user.Username

		u := db.Users[user.Username]

		if u.IsAdmin {
			log.Println("In user.IsAdmin")
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}

	if err := render.Template(w, r, "login.page.html",
		&render.TemplateData{
			Data: make(map[string]interface{}),
			Form: form.New(nil)}); err != nil {
		log.Println("Login: ", err)
		return
	}
}
