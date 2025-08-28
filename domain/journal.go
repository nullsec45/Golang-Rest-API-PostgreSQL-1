package domain 

import (
	"context"
	"database/sql"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
)

const (
	JOURNAL_STATUS_IN_PROGRESS="IN_PROGRESS"
	JOURNAL_STATUS_COMPLETED="COMPLETED"
)

type Journal struct {
	Id string `db:"id"`
	BookId string `db:"book_id"`
	StockCode string `db:"stock_code"`
	CustomerId string `db:"customer_id"`
	Status string `db:"status"`
	DueAt sql.NullTime `db:"due_at"`
	BorrowedAt sql.NullTime `db:"borrowed_at"`
	ReturnedAt sql.NullTime `db:"returned_at"`
}

type JournalSearch struct {
	CustomerId string
	Status string
}

type JournalRepository interface {
	Find(ctx context.Context, se JournalSearch) ([]Journal, error)
	FindById(ctx context.Context, id string) (Journal, error)
	Save(ctx context.Context, journal *Journal) error
	Update(ctx context.Context, journal *Journal) error
}

type JournalService interface {
	Index(ctx context.Context, se JournalSearch) ([]dto.JournalData, error)
	Create(ctx context.Context, req dto.CreateJournalRequest) error
	Return(ctx context.Context, req dto.ReturnJournalRequest) error
}