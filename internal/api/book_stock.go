package api

import (
	"strings"
	"context"
	"net/http"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/utility"
)

type BookStockAPI struct {
	bookStockService domain.BookStockService
} 

func NewBookStock(app *fiber.App, bookStockService domain.BookStockService, authMiddleware fiber.Handler){
	bsa := BookStockAPI {
		bookStockService:bookStockService,
	}

	app.Post("/book-stocks", authMiddleware, bsa.Create)
	app.Delete("/book-stocks", authMiddleware,bsa.Delete)
}

func (ba BookStockAPI) Create (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.CreateBookStockRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := utility.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validasi gagal", fails))
	}
	
	err := ba.bookStockService.Create(c,req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully add stock to book"))
}


func (ba BookStockAPI) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()


	codeStr := ctx.Query("codes")
	if codeStr == "" {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError("parameter code wajib diisi"))
	}

	codes := strings.Split(codeStr,";")
	// if len(codes) < 1 {
	// 	return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError("parameter code wajib diisi"))
	// }

	err := ba.bookStockService.Delete(c, dto.DeleteBookStockRequest{Codes:codes})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully deleted stock book"))
}