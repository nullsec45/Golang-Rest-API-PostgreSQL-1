package service

import (
	"context"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
)

type BookStockService struct {
	bookRepository domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBookStock(bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository) domain.BookStockService {
	return &BookStockService{
		bookRepository: bookRepository,
		bookStockRepository:bookStockRepository,
	}
}

func (b BookStockService) Create(ctx context.Context, req dto.CreateBookStockRequest) error {
	book, err := b.bookRepository.FindById(ctx, req.BookId)

	if err != nil {
		return err
	}

	if book.Id == "" {
		return domain.BookNotFound
	}


	stocks := make([]domain.BookStock, 0)
	for _, v := range req.Codes {
		stocks = append(stocks, domain.BookStock{
			Code:v,
			BookId:req.BookId,
			Status:domain.BOOK_STOCK_STATUS_AVAILABLE,
		})
	}

	return b.bookStockRepository.Save(ctx, stocks)
}

func (b BookStockService) Delete(ctx context.Context, req dto.DeleteBookStockRequest) error {
	return b.bookStockRepository.DeleteByCodes(ctx, req.Codes)
}