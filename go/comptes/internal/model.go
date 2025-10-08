package internal

import (
	"time"

	"github.com/gofrs/uuid"
)

type Budget struct {
	ID     uuid.UUID `json:"id"`
	Label  string    `json:"label"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
	TagID  uuid.UUID `json:"tag_id"`
}

type Income struct {
	ID     int64     `json:"id"`
	Label  string    `json:"label"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
}

type Expense struct {
	ID       uuid.UUID `json:"id"`
	Label    string    `json:"label"`
	Amount   float64   `json:"amount"`
	Date     time.Time `json:"date"`
	BudgetID uuid.UUID `json:"budget_id"`
}

type Tag struct {
	ID          uuid.UUID `json:"id"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
}
