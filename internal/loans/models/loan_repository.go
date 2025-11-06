package models

type LoanRepository interface {
	CreateLoan(loan *Loan) error
	ReturnBook(loan *Loan) error
	GetLoan(id int64) (*Loan, error)
	GetActiveUserLoans(userId int64) ([]*Loan, error)
	GetAllLoans() ([]*Loan, error)
}
