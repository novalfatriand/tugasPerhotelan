package models

type Booking struct {
	ID			int		`json: "id"`
	HotelID		int		`json: "hotel_id"`
	UserID		int 	`json: "user_id"`
	CheckIn		string 	`json: "check_in"`
	CheckOut	string	`json: "check_out"`
	TotalPrice	int		`json: "total_price"`
	Confirmed	bool 	`json: "confirmed"`
}