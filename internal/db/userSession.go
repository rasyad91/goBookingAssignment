package db

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	CustomerID int
	Username   string
	Password   []byte
	Firstname  string
	Lastname   string
	IsAdmin    bool
	BookingsID []int
}

func InitializeUsers() {
	list := []*User{
		{
			Username: "admin",
			Password: []byte("1234"),
			IsAdmin:  true,
		}, {
			Username:   "user",
			Firstname:  "John",
			Lastname:   "McFee",
			Password:   []byte("1234"),
			IsAdmin:    false,
			BookingsID: []int{1, 2, 3, 4, 5},
		},
	}

	for _, u := range list {
		CreateNewUser(u)
	}
}

var Users = map[string]*User{}     //maps username to User
var Sessions = map[string]string{} //maps cookie value to username

func CustomerID() int {
	max := 0
	for _, v := range Users {
		if v.CustomerID > max {
			max = v.CustomerID
		}
	}
	return max + 1
}

func CreateNewUser(u *User) error {

	u.CustomerID = CustomerID()

	bpassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("CreateNewUser: %w", err)
	}
	u.Password = bpassword
	Users[u.Username] = u
	return nil
}

func (u *User) CancelBooking(id int) error {

	for i, v := range u.BookingsID {
		if v == id {
			u.BookingsID = append(u.BookingsID[:i], u.BookingsID[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("index of booking ID not found")
}

func DeleteBookingsFromUserBooking(venueId int) {

	for _, u := range Users {
		if len(u.BookingsID) == 0 {
			continue
		}
	x:
		for _, v := range u.BookingsID {
			booking := BookingList[v]

			if booking.VenueID == venueId {
				u.CancelBooking(v)
				goto x
			}

		}

	}
}
