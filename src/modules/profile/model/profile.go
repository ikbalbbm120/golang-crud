package model

import (
	"time"
)

//profile

type profile struct {
	ID	string		`bson:"id"`
	FirstName string	`bson:"first_name"`
	LastName  string    `bson:"last_name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

//profiles

type profiles []profile