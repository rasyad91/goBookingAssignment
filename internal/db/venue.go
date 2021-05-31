package db

import "fmt"

type Venue struct {
	ID       int
	Name     string
	Capacity int
	Type     string
}

var VenueList []*Venue

func InitializeVenues() {
	VenueList = []*Venue{
		{1, "Undercove", 30, "Hall"},
		{2, "Foundry", 50, "Event Space"},
		{3, "Lunarly", 10, "Meeting Room"},
	}
}

// incrementID helps to increment ID when adding a new Venue
func incrementID() int {
	max := 0
	for _, venue := range VenueList {
		if venue.ID > max {
			max = venue.ID
		}
	}
	return max + 1
}

func GetByID(ID int) *Venue {
	for _, v := range VenueList {
		if v.ID == ID {
			return v
		}
	}
	return nil
}

func Add(v *Venue) {
	v.ID = incrementID()
	VenueList = append(VenueList, v)
}

func DeleteVenue(id int) error {
	for i, v := range VenueList {
		if v.ID == id {
			VenueList = append(VenueList[:i], VenueList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("index of venue not found")
}
