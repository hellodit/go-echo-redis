package domain

import "time"

type (
	Author struct {
		ID int `json:"id"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Email string `json:"email"`
		BirthDate time.Time `json:"birth_date"`
		Added time.Time `json:"added"`
	}
)