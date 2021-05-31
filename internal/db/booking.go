package db

type Booking struct {
	BookingID  int
	CustomerID int
	VenueID    int
	Date       string
}

// BookingList maps BookingID to Booking
var BookingList map[int]Booking
var BookingID int

func init() {
	BookingList = make(map[int]Booking)
	list := []Booking{
		{1, 1, 1, "2021-03-31"},
		{2, 1, 1, "2021-04-03"},
		{3, 1, 1, "2021-04-02"},
		{1, 1, 2, "2021-03-31"},
		{2, 1, 2, "2021-04-03"},
		{3, 1, 2, "2021-04-02"},
	}

	for _, b := range list {
		NewBooking(b)
	}

}

func BookingIsAvailable(venueID int, date string) bool {
	for _, b := range BookingList {
		if b.VenueID == venueID && b.Date == date {
			return false
		}
	}
	return true
}

func NewBooking(b Booking) int {
	BookingID++
	b.BookingID = BookingID
	BookingList[BookingID] = b

	return BookingID
}

func (b *Booking) GetVenueName() string {
	// idInt, _ := strconv.Atoi(id)

	return GetByID(b.VenueID).Name
}

func (b *Booking) GetVenueCapacity() int {
	// idInt, _ := strconv.Atoi(id)

	return GetByID(b.VenueID).Capacity
}

func (b *Booking) GetVenueType() string {
	// idInt, _ := strconv.Atoi(id)

	return GetByID(b.VenueID).Type
}

func DeleteBookingsFromBookingList(id int) error {

	for k, b := range BookingList {
		if b.VenueID == id {
			delete(BookingList, k)
		}
	}
	return nil
}
