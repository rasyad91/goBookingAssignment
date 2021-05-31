package handler

import (
	"assignment-3/internal/db"
	"fmt"
	"net/http"
	"time"
)

func LogOut(w http.ResponseWriter, r *http.Request) {
	if getUser(r) == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")
	start := time.Now()

	delete(db.Sessions, c.Value)

	c = &http.Cookie{
		Name:   "session",
		MaxAge: -1,
		Value:  "",
	}
	http.SetCookie(w, c)

	fmt.Println(time.Since(start))
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
