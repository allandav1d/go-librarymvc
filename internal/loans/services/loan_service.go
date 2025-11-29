package services

import (
	"errors"
	bookService "librarymvc/internal/books/models"
	"librarymvc/internal/loans/models"
	userService "librarymvc/internal/users/models"
	"time"
)

type LoanService struct {
	loanRepository models.LoanRepository
	bookService    bookService.BookService
	userService    userService.UserService
}

func NewLoanService(
	loanRepository models.LoanRepository,
	bookService bookService.BookService,
	userService userService.UserService,
) models.LoanService {
	return &LoanService{
		loanRepository: loanRepository,
		bookService:    bookService,
		userService:    userService,
	}
}

func (l *LoanService) CreateLoan(bookId, userId int64) (*models.Loan, error) {
	book, err := l.bookService.GetBook(bookId)
	if err != nil {
		return nil, err
	}

	// Check if book can be borrowed
	if book.BookType == "referencia" {
		return nil, errors.New("livro de referência não pode ser emprestado - deve permanecer na biblioteca")
	}

	if book.Quantity <= 0 {
		return nil, errors.New("book is not available")
	}

	_, err = l.userService.GetUser(userId)
	if err != nil {
		return nil, err
	}

	activeLoans, err := l.loanRepository.GetActiveUserLoans(userId)
	if err != nil {
		return nil, err
	}

	if len(activeLoans) > 0 {
		return nil, errors.New("user has active loans")
	}

	// Calculate due date based on book's loan duration
	now := time.Now()
	dueDate := now.AddDate(0, 0, book.LoanDuration)

	loan := &models.Loan{
		BookID:     bookId,
		UserID:     userId,
		BorrowedAt: now,
		DueDate:    dueDate,
		Fine:       0,
		Status:     "active",
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	err = l.loanRepository.CreateLoan(loan)
	if err != nil {
		return nil, err
	}

	book.Quantity--
	if err = l.bookService.UpdateBook(book.ID, book); err != nil {
		return nil, err
	}

	return loan, err
}

func (l *LoanService) ReturnBook(loanId int64) error {
	loan, err := l.loanRepository.GetLoan(loanId)
	if err != nil {
		return err
	}

	if loan.Status == "returned" {
		return errors.New("book already returned")
	}

	now := time.Now()
	loan.Status = "returned"
	loan.UpdatedAt = now
	loan.ReturnedAt = now

	// Calculate fine if overdue (R$ 2.00 per day)
	if now.After(loan.DueDate) {
		daysLate := int(now.Sub(loan.DueDate).Hours() / 24)
		if daysLate > 0 {
			loan.Fine = float64(daysLate) * 2.00 // R$ 2.00 por dia de atraso
		}
	}

	if err := l.loanRepository.UpdateLoan(loan); err != nil {
		return err
	}

	book, err := l.bookService.GetBook(loan.BookID)
	if err != nil {
		return err
	}

	book.Quantity++
	return l.bookService.UpdateBook(book.ID, book)
}

// CalculateFine calculates the fine for a given loan
func (l *LoanService) CalculateFine(loan *models.Loan) float64 {
	if loan.Status == "returned" {
		return loan.Fine
	}

	now := time.Now()
	if now.After(loan.DueDate) {
		daysLate := int(now.Sub(loan.DueDate).Hours() / 24)
		if daysLate > 0 {
			return float64(daysLate) * 2.00
		}
	}
	return 0
}

func (l *LoanService) GetLoan(id int64) (*models.Loan, error) {
	return l.loanRepository.GetLoan(id)
}

func (l *LoanService) GetUserLoans(userId int64) ([]*models.Loan, error) {
	return l.loanRepository.GetActiveUserLoans(userId)
}

func (l *LoanService) GetAllLoans() ([]*models.Loan, error) {
	return l.loanRepository.GetAllLoans()
}
