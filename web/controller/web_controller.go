package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	bookModel "librarymvc/internal/books/models"
	loanModel "librarymvc/internal/loans/models"
	userModel "librarymvc/internal/users/models"
)

type WebController struct {
	bookService bookModel.BookService
	userService userModel.UserService
	loanService loanModel.LoanService
}

type DashboardStats struct {
	TotalBooks     int
	TotalUsers     int
	TotalLoans     int
	ActiveLoans    int
	AvailableBooks int
}

type PageData struct {
	Title         string
	ActiveSection string
	FlashMessage  string
	FlashType     string
	Stats         *DashboardStats
	Books         []*bookModel.Book
	Users         []*userModel.User
	Loans         []*loanModel.Loan
	Book          *bookModel.Book
	User          *userModel.User
	Loan          *loanModel.Loan
	IsEdit        bool
	ShowUserLoans bool
	SearchQuery   string
	StatusFilter  string
	BooksMap      map[int64]*bookModel.Book
}

func NewWebController(
	bookService bookModel.BookService,
	userService userModel.UserService,
	loanService loanModel.LoanService,
) *WebController {
	return &WebController{
		bookService: bookService,
		userService: userService,
		loanService: loanService,
	}
}

func (wc *WebController) RegisterRoutes(r *gin.Engine) {
	// Configurar templates
	r.LoadHTMLGlob("templates/*.html")

	// Servir arquivos estáticos
	r.Static("/static", "./static")

	// Rotas principais
	r.GET("/", wc.Dashboard)
	r.GET("/books", wc.BooksList)
	r.GET("/users", wc.UsersList)
	r.GET("/loans", wc.LoansList)

	// Rotas de livros
	r.GET("/books/search", wc.BooksSearch)
	r.GET("/books/:id/edit", wc.BookEditForm)
	r.POST("/books/:id/edit", wc.BookUpdate)
	r.POST("/books/:id/delete", wc.BookDelete)
	r.POST("/books", wc.BookCreate)
	r.POST("/books/create", wc.BookCreate)

	// Rotas de usuários
	r.GET("/users/search", wc.UsersSearch)
	r.GET("/users/:id/edit", wc.UserEditForm)
	r.GET("/users/:id/loans", wc.UserLoans)
	r.POST("/users/:id/edit", wc.UserUpdate)
	r.POST("/users/:id/delete", wc.UserDelete)
	r.POST("/users", wc.UserCreate)
	r.POST("/users/create", wc.UserCreate)

	// Rotas de empréstimos
	r.GET("/loans/search", wc.LoansSearch)
	r.POST("/loans/:id/return", wc.LoanReturn)
	r.POST("/loans", wc.LoanCreate)
	r.POST("/loans/create", wc.LoanCreate)
}

func (wc *WebController) renderTemplate(c *gin.Context, template string, data PageData) {
	c.HTML(http.StatusOK, "layout", data)
}

func (wc *WebController) setFlash(c *gin.Context, message string, flashType string) {
	c.SetCookie("flash_message", message, 3600, "/", "", false, true)
	c.SetCookie("flash_type", flashType, 3600, "/", "", false, true)
}

func (wc *WebController) getFlash(c *gin.Context) (string, string) {
	message, _ := c.Cookie("flash_message")
	flashType, _ := c.Cookie("flash_type")

	// Limpar cookies
	c.SetCookie("flash_message", "", -1, "/", "", false, true)
	c.SetCookie("flash_type", "", -1, "/", "", false, true)

	return message, flashType
}

// Dashboard
func (wc *WebController) Dashboard(c *gin.Context) {
	books, _ := wc.bookService.GetAllBooks()
	users, _ := wc.userService.GetAllUsers()
	loans, _ := wc.loanService.GetAllLoans()

	stats := &DashboardStats{
		TotalBooks:     len(books),
		TotalUsers:     len(users),
		TotalLoans:     len(loans),
		ActiveLoans:    0,
		AvailableBooks: 0,
	}

	for _, loan := range loans {
		if loan.Status == "active" {
			stats.ActiveLoans++
		}
	}

	for _, book := range books {
		if book.Quantity > 0 {
			stats.AvailableBooks++
		}
	}

	// Limitar livros e empréstimos para o dashboard
	var availableBooks []*bookModel.Book
	for _, book := range books {
		if book.Quantity > 0 && len(availableBooks) < 5 {
			availableBooks = append(availableBooks, book)
		}
	}

	var activeLoans []*loanModel.Loan
	for _, loan := range loans {
		if loan.Status == "active" && len(activeLoans) < 5 {
			activeLoans = append(activeLoans, loan)
		}
	}

	var limitedUsers []*userModel.User
	if len(users) > 5 {
		limitedUsers = users[:5]
	} else {
		limitedUsers = users
	}

	message, flashType := wc.getFlash(c)

	data := PageData{
		Title:         "Dashboard - Sistema de Biblioteca",
		ActiveSection: "dashboard",
		FlashMessage:  message,
		FlashType:     flashType,
		Stats:         stats,
		Books:         availableBooks,
		Users:         limitedUsers,
		Loans:         activeLoans,
	}

	wc.renderTemplate(c, "dashboard", data)
}

// Books
func (wc *WebController) BooksList(c *gin.Context) {
	books, err := wc.bookService.GetAllBooks()
	if err != nil {
		books = []*bookModel.Book{}
	}

	message, flashType := wc.getFlash(c)

	data := PageData{
		Title:         "Livros - Sistema de Biblioteca",
		ActiveSection: "books",
		FlashMessage:  message,
		FlashType:     flashType,
		Books:         books,
	}

	wc.renderTemplate(c, "books", data)
}

func (wc *WebController) BooksSearch(c *gin.Context) {
	query := strings.ToLower(c.Query("q"))
	books, err := wc.bookService.GetAllBooks()
	if err != nil {
		books = []*bookModel.Book{}
	}

	if query != "" {
		var filtered []*bookModel.Book
		for _, book := range books {
			if strings.Contains(strings.ToLower(book.Title), query) ||
				strings.Contains(strings.ToLower(book.Author), query) {
				filtered = append(filtered, book)
			}
		}
		books = filtered
	}

	message, flashType := wc.getFlash(c)

	data := PageData{
		Title:         "Livros - Sistema de Biblioteca",
		ActiveSection: "books",
		FlashMessage:  message,
		FlashType:     flashType,
		Books:         books,
		SearchQuery:   query,
	}

	wc.renderTemplate(c, "books", data)
}

func (wc *WebController) BookEditForm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		wc.setFlash(c, "ID inválido", "error")
		c.Redirect(http.StatusFound, "/books")
		return
	}

	book, err := wc.bookService.GetBook(id)
	if err != nil {
		wc.setFlash(c, "Livro não encontrado", "error")
		c.Redirect(http.StatusFound, "/books")
		return
	}

	message, flashType := wc.getFlash(c)

	data := PageData{
		Title:         "Editar Livro - Sistema de Biblioteca",
		ActiveSection: "books",
		FlashMessage:  message,
		FlashType:     flashType,
		Book:          book,
		IsEdit:        true,
	}

	wc.renderTemplate(c, "books", data)
}

func (wc *WebController) BookUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		wc.setFlash(c, "ID inválido", "error")
		c.Redirect(http.StatusFound, "/books")
		return
	}

	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	book := &bookModel.Book{
		Title:    c.PostForm("title"),
		Author:   c.PostForm("author"),
		Quantity: quantity,
	}

	err = wc.bookService.UpdateBook(id, book)
	if err != nil {
		wc.setFlash(c, "Erro ao atualizar livro: "+err.Error(), "error")
		c.Redirect(http.StatusFound, "/books/"+c.Param("id")+"/edit")
		return
	}

	wc.setFlash(c, "Livro atualizado com sucesso!", "success")
	c.Redirect(http.StatusFound, "/books")
}

func (wc *WebController) BookDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		wc.setFlash(c, "ID inválido", "error")
		c.Redirect(http.StatusFound, "/books")
		return
	}

	err = wc.bookService.DeleteBook(id)
	if err != nil {
		wc.setFlash(c, "Erro ao excluir livro: "+err.Error(), "error")
	} else {
		wc.setFlash(c, "Livro excluído com sucesso!", "success")
	}

	c.Redirect(http.StatusFound, "/books")
}

func (wc *WebController) BookCreate(c *gin.Context) {
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	book := &bookModel.Book{
		Title:    c.PostForm("title"),
		Author:   c.PostForm("author"),
		Quantity: quantity,
	}

	err := wc.bookService.CreateBook(book)
	if err != nil {
		wc.setFlash(c, "Erro ao criar livro: "+err.Error(), "error")
	} else {
		wc.setFlash(c, "Livro criado com sucesso!", "success")
	}

	c.Redirect(http.StatusFound, "/books")
}

// Users
func (wc *WebController) UsersList(c *gin.Context) {
	users, err := wc.userService.GetAllUsers()
	if err != nil {
		users = []*userModel.User{}
	}

	message, flashType := wc.getFlash(c)

	data := PageData{
		Title:         "Usuários - Sistema de Biblioteca",
		ActiveSection: "users",
		FlashMessage:  message,
		FlashType:     flashType,
		Users:         users,
	}

	wc.renderTemplate(c, "users", data)
}

func (wc *WebController) UsersSearch(c *gin.Context) {
	query := strings.ToLower(c.Query("q"))
	users, err := wc.userService.GetAllUsers()
	if err != nil {
		users = []*userModel.User{}
	}

	if query != "" {
		var filtered []*userModel.User
		for _, user := range users {
			if strings.Contains(strings.ToLower(user.Name), query) ||
				strings.Contains(strings.ToLower(user.Email), query) {
				filtered = append(filtered, user)
			}
		}
		users = filtered
	}

	message, flashType := wc.getFlash(c)

	data := PageData{
		Title:         "Usuários - Sistema de Biblioteca",
		ActiveSection: "users",
		FlashMessage:  message,
		FlashType:     flashType,
		Users:         users,
		SearchQuery:   query,
	}

	wc.renderTemplate(c, "users", data)
}

func (wc *WebController) UserEditForm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		wc.setFlash(c, "ID inválido", "error")
		c.Redirect(http.StatusFound, "/users")
		return
	}

	user, err := wc.userService.GetUser(id)
	if err != nil {
		wc.setFlash(c, "Usuário não encontrado", "error")
		c.Redirect(http.StatusFound, "/users")
		return
	}

	message, flashType := wc.getFlash(c)

	data := PageData{
		Title:         "Editar Usuário - Sistema de Biblioteca",
		ActiveSection: "users",
		FlashMessage:  message,
		FlashType:     flashType,
		User:          user,
		IsEdit:        true,
	}

	wc.renderTemplate(c, "users", data)
}

func (wc *WebController) UserUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		wc.setFlash(c, "ID inválido", "error")
		c.Redirect(http.StatusFound, "/users")
		return
	}

	user := &userModel.User{
		Name:  c.PostForm("name"),
		Email: c.PostForm("email"),
	}

	err = wc.userService.UpdateUser(id, user)
	if err != nil {
		wc.setFlash(c, "Erro ao atualizar usuário: "+err.Error(), "error")
		c.Redirect(http.StatusFound, "/users/"+c.Param("id")+"/edit")
		return
	}

	wc.setFlash(c, "Usuário atualizado com sucesso!", "success")
	c.Redirect(http.StatusFound, "/users")
}

func (wc *WebController) UserDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		wc.setFlash(c, "ID inválido", "error")
		c.Redirect(http.StatusFound, "/users")
		return
	}

	err = wc.userService.DeleteUser(id)
	if err != nil {
		wc.setFlash(c, "Erro ao excluir usuário: "+err.Error(), "error")
	} else {
		wc.setFlash(c, "Usuário excluído com sucesso!", "success")
	}

	c.Redirect(http.StatusFound, "/users")
}

func (wc *WebController) UserCreate(c *gin.Context) {
	user := &userModel.User{
		Name:  c.PostForm("name"),
		Email: c.PostForm("email"),
	}

	err := wc.userService.CreateUser(user)
	if err != nil {
		wc.setFlash(c, "Erro ao criar usuário: "+err.Error(), "error")
	} else {
		wc.setFlash(c, "Usuário criado com sucesso!", "success")
	}

	c.Redirect(http.StatusFound, "/users")
}

func (wc *WebController) UserLoans(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		wc.setFlash(c, "ID inválido", "error")
		c.Redirect(http.StatusFound, "/users")
		return
	}

	user, err := wc.userService.GetUser(id)
	if err != nil {
		wc.setFlash(c, "Usuário não encontrado", "error")
		c.Redirect(http.StatusFound, "/users")
		return
	}

	// Buscar todos os empréstimos e filtrar por usuário
	allLoans, err := wc.loanService.GetAllLoans()
	if err != nil {
		allLoans = []*loanModel.Loan{}
	}

	var loans []*loanModel.Loan
	for _, loan := range allLoans {
		if loan.UserID == id {
			loans = append(loans, loan)
		}
	}

	// Buscar todos os livros para criar o mapa
	allBooks, _ := wc.bookService.GetAllBooks()
	booksMap := make(map[int64]*bookModel.Book)
	for _, book := range allBooks {
		booksMap[book.ID] = book
	}

	message, flashType := wc.getFlash(c)

	data := PageData{
		Title:         "Empréstimos do Usuário - Sistema de Biblioteca",
		ActiveSection: "users",
		FlashMessage:  message,
		FlashType:     flashType,
		User:          user,
		Loans:         loans,
		ShowUserLoans: true,
		BooksMap:      booksMap,
	}

	wc.renderTemplate(c, "users", data)
}

// Loans
func (wc *WebController) LoansList(c *gin.Context) {
	loans, err := wc.loanService.GetAllLoans()
	if err != nil {
		loans = []*loanModel.Loan{}
	}

	message, flashType := wc.getFlash(c)

	data := PageData{
		Title:         "Empréstimos - Sistema de Biblioteca",
		ActiveSection: "loans",
		FlashMessage:  message,
		FlashType:     flashType,
		Loans:         loans,
	}

	wc.renderTemplate(c, "loans", data)
}

func (wc *WebController) LoansSearch(c *gin.Context) {
	query := strings.ToLower(c.Query("q"))
	statusFilter := c.Query("status")
	loans, err := wc.loanService.GetAllLoans()
	if err != nil {
		loans = []*loanModel.Loan{}
	}

	var filtered []*loanModel.Loan
	for _, loan := range loans {
		// Filtro por status
		if statusFilter != "" && loan.Status != statusFilter {
			continue
		}

		// Busca por texto (pode buscar por ID do livro ou usuário)
		if query != "" {
			bookIdStr := strconv.FormatInt(loan.BookID, 10)
			userIdStr := strconv.FormatInt(loan.UserID, 10)
			if !strings.Contains(bookIdStr, query) && !strings.Contains(userIdStr, query) {
				continue
			}
		}

		filtered = append(filtered, loan)
	}

	message, flashType := wc.getFlash(c)

	data := PageData{
		Title:         "Empréstimos - Sistema de Biblioteca",
		ActiveSection: "loans",
		FlashMessage:  message,
		FlashType:     flashType,
		Loans:         filtered,
		SearchQuery:   query,
		StatusFilter:  statusFilter,
	}

	wc.renderTemplate(c, "loans", data)
}

func (wc *WebController) LoanReturn(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		wc.setFlash(c, "ID inválido", "error")
		c.Redirect(http.StatusFound, "/loans")
		return
	}

	err = wc.loanService.ReturnBook(id)
	if err != nil {
		wc.setFlash(c, "Erro ao devolver livro: "+err.Error(), "error")
	} else {
		wc.setFlash(c, "Livro devolvido com sucesso!", "success")
	}

	c.Redirect(http.StatusFound, "/loans")
}

func (wc *WebController) LoanCreate(c *gin.Context) {
	bookId, err := strconv.ParseInt(c.PostForm("book_id"), 10, 64)
	if err != nil {
		wc.setFlash(c, "ID do livro inválido", "error")
		c.Redirect(http.StatusFound, "/loans")
		return
	}

	userId, err := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	if err != nil {
		wc.setFlash(c, "ID do usuário inválido", "error")
		c.Redirect(http.StatusFound, "/loans")
		return
	}

	_, err = wc.loanService.CreateLoan(bookId, userId)
	if err != nil {
		wc.setFlash(c, "Erro ao criar empréstimo: "+err.Error(), "error")
	} else {
		wc.setFlash(c, "Empréstimo criado com sucesso!", "success")
	}

	c.Redirect(http.StatusFound, "/loans")
}
