package api

import (
	"net/http"
	"context"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/utility"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
)

type JournalAPI struct {
	journalService domain.JournalService
}

func NewJournal(app *fiber.App, journalService domain.JournalService, authMiddleware fiber.Handler){
	ja := JournalAPI {
		journalService:journalService,
	}

	app.Get("/journals", authMiddleware, ja.Index)
	app.Post("/journals", authMiddleware, ja.Create)
	app.Put("/journals/:id", authMiddleware, ja.Update)
}

func (ja JournalAPI) Index (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	customerId :=  ctx.Query("customer_id")
	status := ctx.Query("status")
	res, err :=	ja.journalService.Index(c, domain.JournalSearch{
		CustomerId:customerId,
		Status:status,
	}) 

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfullly Get Data",res))
}

func (ja JournalAPI) Create (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.CreateJournalRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := utility.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("Validate Error", fails))
	}

	err :=	ja.journalService.Create(c, req) 

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfullly Created Data"))
}


func (ja JournalAPI) Update (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")

	err :=	ja.journalService.Return(c, dto.ReturnJournalRequest{
		JournalId:id,
	}) 

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfullly Updated Data"))
}
