package models

import "time"

type Book struct {
	ID           int64     `json:"ID"`
	Title        string    `json:"title" binding:"required,min=5"`
	Author       string    `json:"author" binding:"required,min=5"`
	Quantity     int       `json:"quantity" binding:"required,min=0"`
	BookType     string    `json:"bookType" binding:"required,oneof=emprestavel referencia"` // emprestavel or referencia
	LoanDuration int       `json:"loanDuration"`                                             // 6, 12, or 30 days (only for emprestavel)
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
