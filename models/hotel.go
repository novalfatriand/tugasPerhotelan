package models

type Hotel struct {
	ID				int		`json: "id"`
	Name			string	`json: "name"`
	Location		string 	`json: "location"`
	PricePerNight	int 	`json: "price_per_night"`
}