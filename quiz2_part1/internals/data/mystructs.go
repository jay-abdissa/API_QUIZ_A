// Filename: internal/data/mystructs.go

package data

import (
	"time"
)

type Mystruct struct {
	ID int64 `json:"id"`
	CreatedAt time.Time `json:"createdat"`
	Name string `json:"name"`
	Year string `json:"year"`
	Contact string `json:"contact"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Website string `json:"website"`
	Address string `json:"address"`
}