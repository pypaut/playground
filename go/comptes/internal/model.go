package internal

import (
	"time"
)

type Budget struct {
	Label  string    `json:"label"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
	Tag    string    `json:"tag"`
}

type Income struct {
	Label  string    `json:"label"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
}

type Expense struct {
	Label  string    `json:"label"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
	Budget string    `json:"budget"`
}

type Tag struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
