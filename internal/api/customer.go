package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
	"time"
	"net/http"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/utility"
	// "fmt"
)

type customerAPI struct {
	customerService domain.CustomerService
}

func NewCustomer(app * fiber.App, customerService domain.CustomerService) {
	ca := customerAPI{
		customerService: customerService,
	}

	app.Get("/customers", ca.Index)
	app.Post("/customers", ca.Create)
}

func (ca customerAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	res, err := ca.customerService.Index(c)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Get Data",res))
}

func (ca customerAPI) Create (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := utility.Validate(req)
	
	if len(fails) > 0{
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData(
			"validation failed",
			fails,
		))
	}

	res, err := ca.customerService.Create(c, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully Created Data", res))
}