package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
	"time"
)

type customerAPI struct {
	customerService domain.CustomerService
}

func NewCustomer(app * fiber.App, customerService domain.CustomerService) {
	ca := customerAPI{
		customerService: customerService,
	}

	app.Get("/customers", ca.Index)
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