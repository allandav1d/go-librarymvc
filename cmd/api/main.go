package main

import (
	"github.com/gin-gonic/gin"
	usercontroller "librarymvc/internal/users/controllers"
	bookcontroller "librarymvc/internal/books/controllers"
	loancontroller "librarymvc/internal/loans/controllers"
	"log"
)

func main() {
	router := gin.Default()

	booksController := bookcontroller.NewBooksController()
	usersController := usercontroller.NewUserController()
	loansControoler := loancontroller.NewLoanController()

	booksController.RegisterRoutes(router)
	usersController.RegisterRoutes(router)
	loansControoler.RegisterRoutes(router)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
