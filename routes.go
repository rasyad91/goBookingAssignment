package main

import (
	"assignment-3/internal/handler"
	"net/http"
)

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", handler.Home)
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/login", handler.LogIn)
	mux.HandleFunc("/admin", handler.Admin)
	mux.HandleFunc("/admin/venue", handler.Venue)

	mux.HandleFunc("/logout", handler.LogOut)
	mux.HandleFunc("/bookings", handler.Bookings)
	mux.HandleFunc("/", handler.NotFound)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
