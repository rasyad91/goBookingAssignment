package handler

import (
	"assignment-3/internal/db"
	"assignment-3/internal/render"
	"log"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	var user db.User
	if o := getUser(r); o != nil {
		user = *o
	}
	data := make(map[string]interface{})
	data["user"] = user
	switch r.URL.Path {
	case "/":
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	default:
		if err := render.Template(w, r, "notfound.page.html", &render.TemplateData{
			Data: data,
		}); err != nil {
			log.Println("Notfound: Error parsing template: ", err)
		}
	}
}
