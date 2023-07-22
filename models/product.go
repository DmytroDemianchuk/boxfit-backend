package models

import "time"

type Product struct {
	ID          string
	UserID      string
	Name        string
	Category    string
	Subcategory string
	Mark        string
	Variant     string
	Color       string
	Number      uint16
	Price       uint32
	Image       []byte
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
