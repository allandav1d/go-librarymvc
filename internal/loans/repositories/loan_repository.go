package repositories

import (
	"errors"
	"librarymvc/internal/loans/models"
	"sync"
)

type LoanRepository struct {
	loans  map[int64]*models.Loan
	mu     sync.RWMutex
	nextID int64
}

func NewLoanRepository() models.LoanRepository {
	return &LoanRepository{
		loans:  make(map[int64]*models.Loan),
		nextID: 1,
	}
}

func (l *LoanRepository) CreateLoan(loan *models.Loan) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan.ID = l.nextID
	l.loans[l.nextID] = loan
	l.nextID++

	return nil
}

func (l *LoanRepository) UpdateLoan(loan *models.Loan) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	_, exists := l.loans[loan.ID]
	if !exists {
		return errors.New("loan not found")
	}
	l.loans[loan.ID] = loan
	return nil
}

func (l *LoanRepository) ReturnBook(loan *models.Loan) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, exists := l.loans[loan.ID]
	if !exists {
		return errors.New("loan not found")
	}

	loan.Status = "returned"
	l.loans[loan.ID] = loan

	return nil
}

func (l *LoanRepository) GetLoan(id int64) (*models.Loan, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	loan, exists := l.loans[id]
	if !exists {
		return nil, errors.New("loan not found")
	}

	return loan, nil
}

func (l *LoanRepository) GetActiveUserLoans(userId int64) ([]*models.Loan, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	activeLoans := make([]*models.Loan, 0)
	for _, loan := range l.loans {
		if loan.UserID == userId && loan.Status == "active" {
			activeLoans = append(activeLoans, loan)
		}
	}

	return activeLoans, nil
}

func (l *LoanRepository) GetAllLoans() ([]*models.Loan, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	loans := make([]*models.Loan, 0, len(l.loans))
	for _, loan := range l.loans {
		loans = append(loans, loan)
	}

	return loans, nil
}
