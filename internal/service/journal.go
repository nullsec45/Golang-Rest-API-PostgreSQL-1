package service

import (
	"context"
	"time"
    "github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
	"github.com/google/uuid"
	"database/sql"
	"fmt"
	"errors"
)

type JournalService struct {
	journalRepository domain.JournalRepository
	bookRepository domain.BookRepository
	bookStockRepository domain.BookStockRepository
	customerRepository domain.CustomerRepository
	chargeRepository domain.ChargeRepository
}

func NewJournal(
		journalRepository domain.JournalRepository,
		bookRepository domain.BookRepository,
		bookStockRepository domain.BookStockRepository,
		customerRepository domain.CustomerRepository, 
		chargeRepository domain.ChargeRepository,
	) domain.JournalService {
		return &JournalService{
			journalRepository:journalRepository,
			bookRepository:bookRepository,
			bookStockRepository:bookStockRepository,
			customerRepository:customerRepository,
			chargeRepository:chargeRepository,
		}
	}

func (j JournalService) Index (ctx context.Context, se domain.JournalSearch) ([]dto.JournalData, error) {
	journals, err := j.journalRepository.Find(ctx, se)
	if err != nil {
		return nil, err
	}
	customerId := make([]string, 0)
	bookId := make([]string, 0)
	for _, v := range journals {
		customerId = append(customerId, v.CustomerId)
		bookId=append(bookId, v.BookId)
	}
	customers := make(map[string]domain.Customer)
	if len(customerId) > 0 {
		customersDb, _ := j.customerRepository.FindByIds(ctx, customerId)
		
		for _, v := range customersDb {
			customers[v.ID] = v
		}
	}

	books := make(map[string]domain.Book)

	if len(bookId) > 0 {
		bookDb, _ := j.bookRepository.FindByIds(ctx, bookId)

		for _, v := range bookDb {
			books[v.Id] = v
		}
	}

	result := make([]dto.JournalData, 0)
	for _, v := range journals {
		book := dto.BookData{}
		if v2, e := books[v.BookId]; e {
			book = dto.BookData {
				Id:v2.Id,
				Isbn: v2.Isbn,
				Title: v2.Title,
				Description : v2.Description,
			}
		}

		customer := dto.CustomerData{}
		if v2, e := customers[v.CustomerId]; e {
			customer = dto.CustomerData{
				ID: v2.ID,
				Code:v2.Code,
				Name:v2.Name,
			}
		}

		result = append(result, dto.JournalData{
			Id:v.Id,
			BookStock:v.StockCode,
			Book:book,
			Customer: customer,
			BorrowedAt:v.BorrowedAt.Time,
			ReturnedAt:v.ReturnedAt.Time,
		})
	}

	return result, nil
}

func (j JournalService) Create (ctx context.Context, req dto.CreateJournalRequest) error {
	book, err := j.bookRepository.FindById(ctx, req.BookId)
	if err != nil {
		return err
	}

	if book.Id == "" {
		return domain.BookNotFound
	}

	stock, err := j.bookStockRepository.FindByBookAndCode(ctx, book.Id, req.BookStock)

	fmt.Println(err)

	if err != nil {
		return err
	}

	if stock.Code == "" {
		return domain.BookNotFound
	}

	if stock.Status != domain.BOOK_STOCK_STATUS_AVAILABLE {
		return errors.New("stock buku sudah dipinjam sebelumnya")
	}

	journal := domain.Journal{
		Id:uuid.NewString(),
		BookId:req.BookId,
		StockCode:req.BookStock,
		CustomerId:req.CustomerId,
		Status:domain.JOURNAL_STATUS_IN_PROGRESS,
		DueAt:sql.NullTime{Valid:true, Time:time.Now().Add(7 * 24 * time.Hour)},
		BorrowedAt:sql.NullTime{Valid:true, Time:time.Now()},
	}

	err = j.journalRepository.Save(ctx, &journal)

	if err != nil {
		return err
	}

	stock.Status=domain.BOOK_STOCK_STATUS_BORROWED
	stock.BorrowedAt=journal.BorrowedAt
	stock.BorrowerId=sql.NullString{Valid:true, String:journal.CustomerId}
	return j.bookStockRepository.Update(ctx, &stock)
}

func (j JournalService) Return (ctx context.Context, req dto.ReturnJournalRequest) error {
	journal, err := j.journalRepository.FindById(ctx, req.JournalId)
	if err != nil {
		return err
	}

	if journal.Id == "" {
		return domain.JournalNotFound
	}

	stock, err := j.bookStockRepository.FindByBookAndCode(ctx,journal.BookId, journal.StockCode)

	if err != nil {
		return err
	}

	if stock.Code != "" {
		stock.Status=domain.BOOK_STOCK_STATUS_AVAILABLE
		stock.BorrowerId = sql.NullString{Valid:false}
		stock.BorrowedAt = sql.NullTime{Valid:false}
		err = j.bookStockRepository.Update(ctx, &stock)
		if err != nil {
			return err
		}
	}

	journal.Status=domain.JOURNAL_STATUS_COMPLETED
	journal.ReturnedAt=sql.NullTime{Valid:true, Time:time.Now()}
	err = j.journalRepository.Update(ctx, &journal)

	if err != nil {
		return err
	}

	hoursLate := time.Now().Sub(journal.DueAt.Time).Hours()
	if hoursLate > 24 {
		daysLate := int(hoursLate / 24)
		charge := domain.Charge{
			Id:uuid.NewString(),
			JournalId:journal.Id,
			DaysLate:daysLate,
			DailyLateFee:5000,
			Total: 5000 * daysLate,
			UserId: req.UserId,
		}

		err = j.chargeRepository.Save(ctx, &charge)
	}

	return err
}