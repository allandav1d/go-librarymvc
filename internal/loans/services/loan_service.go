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

	loan := &models.Loan{
		BookID:     bookId,
		UserID:     userId,
		BorrowedAt: time.Now(),
		Status:     "active",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
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

	loan.Status = "returned"
	loan.UpdatedAt = time.Now()
	loan.ReturnedAt = time.Now()

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

func (l *LoanService) GetLoan(id int64) (*models.Loan, error) {
	return l.loanRepository.GetLoan(id)
}

func (l *LoanService) GetUserLoans(userId int64) ([]*models.Loan, error) {
	return l.loanRepository.GetActiveUserLoans(userId)
}

func (l *LoanService) GetAllLoans() ([]*models.Loan, error) {
	return l.loanRepository.GetAllLoans()
}
