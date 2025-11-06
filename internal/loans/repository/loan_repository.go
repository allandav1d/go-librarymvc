package repository

import (
	"library/internal/loans/models"
	"sync"
	"errors"
)

type LoanRepository struct {
	loans   map[int64]*models.Loan
	mu      sync.RWMutex
	nextID  int64
}

func (l LoanRepository) CreateLoan(loan *models.Loan) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan.ID = l.nextID
	l.loans[l.nextID] = loan
	l.nextID++

	return nil
}

func (l LoanRepository) UpdateLoan(loan *models.Loan) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	_, exists := l.loans[loan.ID]
	if !exists {
		return errors.New("loan not found")
	}
	l.loans[loan.ID] = loan
	return nil
}

func (l LoanRepository) ReturnBook(loanId int64) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan, exists := l.loans[loanId]
	if !exists {
		return errors.New("loan not found")
	}
	delete(l.loans, loanId)

	loan.Status = models.LoanStatusReturned

	return nil
}

func (l *LoanRepository) GetLoan(id int64) (*models.Loan, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan, exists := l.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}

	return loan, nil
}

func (l LoanRepository) GetActiveUserLoans(userId int64) ([]*models.Loan, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	activeLoans := make([]*models.Loan, 0)
	for _, loan := range l.loans {
		if loan.UserID == userId && loan.Status == models.LoanStatusActive {
			activeLoans = append(activeLoans, loan)
		}
	}

	if len(activeLoans) == 0 {
		return nil, errors.New("no active loans found for user")
	}

	return activeLoans, nil
}

func (l LoanRepository) GetAllLoans() ([]*models.Loan, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	loans := make([]*models.Loan, 0, len(l.loans))
	for _, loan := range l.loans {
		loans = append(loans, loan)
	}

	return loans, nil
}

