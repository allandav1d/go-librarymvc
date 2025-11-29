package models

import "time"

type Loan struct {
	ID         int64     `json:"ID"`
	BookID     int64     `json:"bookID"`
	UserID     int64     `json:"userID"`
	BorrowedAt time.Time `json:"borrowedAt"`
	DueDate    time.Time `json:"dueDate"`    // Data de devolução prevista
	ReturnedAt time.Time `json:"returnedAt"`
	Fine       float64   `json:"fine"`   // Multa por atraso (R$ 2.00 por dia)
	Status     string    `json:"status"` // active, returned, overdue
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
