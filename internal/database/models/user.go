package models

// I know that I could use an ORM to make all these things easier, but an ORM would come with some complexity too!
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
