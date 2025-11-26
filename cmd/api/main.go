package main

import (
	"log"

	"github.com/gin-gonic/gin"

	bookcontroller "librarymvc/internal/books/controllers"
	bookrepository "librarymvc/internal/books/repositories"
	bookservice "librarymvc/internal/books/services"

	usercontroller "librarymvc/internal/users/controllers"
	userrepository "librarymvc/internal/users/repositories"
	userservice "librarymvc/internal/users/services"

	loancontroller "librarymvc/internal/loans/controllers"
	loanrepository "librarymvc/internal/loans/repositories"
	loanservice "librarymvc/internal/loans/services"

	webcontroller "librarymvc/web/controller"
)

func main() {
	router := gin.Default()

	// Initialize repositories
	bookRepo := bookrepository.NewBookRepository()
	userRepo := userrepository.NewUserRepository()
	loanRepo := loanrepository.NewLoanRepository()

	// Initialize services
	bookSvc := bookservice.NewBookService(bookRepo)
	userSvc := userservice.NewUserService(userRepo)
	loanSvc := loanservice.NewLoanService(loanRepo, bookSvc, userSvc)

	// Initialize Web controller
	webController := webcontroller.NewWebController(bookSvc, userSvc, loanSvc)

	// Register Web routes first (they have priority)
	webController.RegisterRoutes(router)

	// Initialize API controllers
	booksController := bookcontroller.NewBooksController(bookSvc)
	usersController := usercontroller.NewUserController(userSvc)
	loansController := loancontroller.NewLoanController(loanSvc)

	// Register API routes with /api prefix
	api := router.Group("/api")
	apiBooks := api.Group("/books")
	{
		apiBooks.GET("/", booksController.GetAllBooks)
		apiBooks.GET("/:id", booksController.GetBook)
		apiBooks.POST("", booksController.CreateBook)
		apiBooks.PUT("/:id", booksController.UpdateBook)
		apiBooks.DELETE("/:id", booksController.DeleteBook)
	}

	apiUsers := api.Group("/users")
	{
		apiUsers.GET("/", usersController.GetAllUsers)
		apiUsers.GET("/:id", usersController.GetUser)
		apiUsers.POST("", usersController.CreateUser)
		apiUsers.PUT("/:id", usersController.UpdateUser)
		apiUsers.DELETE("/:id", usersController.DeleteUser)
	}

	apiLoans := api.Group("/loans")
	{
		apiLoans.POST("", loansController.CreateLoan)
		apiLoans.GET("/:id", loansController.GetLoan)
		apiLoans.GET("", loansController.GetAllLoans)
		apiLoans.PUT("/:id/return", loansController.ReturnBook)
	}

	apiLoansUsers := api.Group("/loans/users")
	{
		apiLoansUsers.GET("/:userId/loans", loansController.GetUserLoans)
	}

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
