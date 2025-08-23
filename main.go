package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/config"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/connection"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/repository"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/service"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/api"
	jwtMid "github.com/gofiber/contrib/jwt"
	"net/http"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
)

func main() {
	// This is the main function where the program starts execution.
	// You can add your code here to implement the desired functionality.
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	authMiddleware := jwtMid.New(
		jwtMid.Config{
			SigningKey:jwtMid.SigningKey{Key:[]byte(cnf.Jwt.Key)},
			ErrorHandler:func (ctx *fiber.Ctx, err error) error {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.CreateResponseError("Endpoint perlu token, silahkan login terlebih dahulu."))
			},
		},
	)

	customerRepository := repository.NewCustomer(dbConnection)
	userRepository := repository.NewUser(dbConnection)
	bookRepository := repository.NewBook(dbConnection)
	bookStockRepository :=  repository.NewBookStock(dbConnection)

	customerService := service.NewCustomer(customerRepository)
	authService := service.NewAuth(cnf, userRepository)
	bookService := service.NewBook(bookRepository, bookStockRepository)
	bookStockService := service.NewBookStock(bookRepository, bookStockRepository)

	api.NewCustomer(app, customerService, authMiddleware)
	api.NewAuth(app, authService)
	api.NewBook(app, bookService, authMiddleware)
	api.NewBookStock(app, bookStockService, authMiddleware)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}

func developers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")
}