package service

import (
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
	"context"
	"github.com/google/uuid"
	"database/sql"
	"time"
	"errors"
    // "fmt"
)

type BookService struct {
	bookRepository domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBook(
	bookRepository domain.BookRepository,
	bookStockRepository domain.BookStockRepository,
) domain.BookService {
	return &BookService{
		bookRepository: bookRepository,
		bookStockRepository: bookStockRepository,
	}
}

func (bs BookService) Index(ctx context.Context)  ([]dto.BookData, error) {
	books, err := bs.bookRepository.FindAll(ctx)
	if err != nil {
		// Handle error (e.g., log it, return an error response)
		return nil, err
	}

	var bookData []dto.BookData

	for _, v := range books {
		bookData = append(bookData, dto.BookData{
			Id:   v.Id,
			Isbn: v.Isbn,
			Title: v.Title,
			Description: v.Description,
		})
	}
	
	return bookData, nil
}

func (bs BookService) Show (ctx context.Context, id string) (dto.BookData, error) {
    data, err := bs.bookRepository.FindById(ctx,id)

    if err != nil && data.Id == "" {
        return dto.BookData{}, errors.New("Data buku tidak ditemukan!.")
    }
    
    if err != nil {
        return dto.BookData{}, err
    }

    return dto.BookData{
		Id:   data.Id,
		Isbn: data.Isbn,
		Title: data.Title,
		Description: data.Description,
    }, nil
}

func (bs BookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
    book := domain.Book{
        Id: uuid.New().String(),
       	Isbn: req.Isbn,
		Title: req.Title,
		Description: req.Description,
		CreatedAt:sql.NullTime{Valid:true, Time:time.Now()},	
    }

	return bs.bookRepository.Save(ctx, &book)
}

func (bs BookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
    // Cari data book
    exist, err := bs.bookRepository.FindById(ctx, req.Id)

    // Jika book tidak ditemukan
    if err != nil && exist.Id == "" {
        return  errors.New("Data buku tidak ditemukan!.")
    }
    
    if err != nil {
        return  err
    }

	exist.Isbn=req.Isbn
	exist.Title=req.Title
	exist.Description=req.Description
	exist.UpdatedAt=sql.NullTime{Valid:true, Time:time.Now()}

    return bs.bookRepository.Update(ctx, &exist)
}

func (bs BookService) Delete (ctx context.Context, id string) error {
    exist, err := bs.bookRepository.FindById(ctx,id)

    if err != nil && exist.Id == "" {
        return  errors.New("Data buku tidak ditemukan!.")
    }
    
    if err != nil {
        return err
    }

	err = bs.bookRepository.Delete(ctx,exist.Id)

	if err != nil {
		return err
	}

    return bs.bookStockRepository.DeleteByBookId(ctx, exist.Id)
}