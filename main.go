package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/config"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/connection"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/repository"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/service"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/api"
)

func main() {
	// This is the main function where the program starts execution.
	// You can add your code here to implement the desired functionality.
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)
	customerService := service.NewCustomer(customerRepository)
	api.NewCustomer(app, customerService)
	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}

func developers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")
}