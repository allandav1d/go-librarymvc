package models

type LoanService interface {
	CreateLoan(bookID, userID int64) (*Loan, error)
	ReturnBook(loanID int64) error
	GetLoan(id int64) (*Loan, error)
	GetUserLoans(userID int64) ([]*Loan, error)
	GetAllLoans() ([]*Loan, error)
}
