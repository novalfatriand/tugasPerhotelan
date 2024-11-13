package models

type User struct {
	ID			int		`json: "id"`
	Username	string	`json: "name"`
	Password	string 	`json: "location"`
	Role		string 	`json: "price_per_night"`
}