package models

import "time"

type WishlistEntry struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Username    string    `json:"username"`
	CountryName string    `json:"name"`
	Slug        string    `json:"slug"`
	Note        string    `json:"note"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
