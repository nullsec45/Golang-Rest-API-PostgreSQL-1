package domain

import (
	"database/sql"
	"context"
)

type BookStock struct {
	Code string `db:"code"`
	BookId string `db:"book_id"`
	Status string `db:"status"`
	BorrowerId sql.NullString `db:"borrower_id"`
	BorrowedAt sql.NullTime `db:"borrowed_at"`
}

type BookStockRepository interface {
	FindByBookId(ctx context.Context, id string) ([]BookStock, error)
	FindByBookAndCode(ctx context.Context, id string, code string) (BookStock, error)
	Save(ctx context.Context, data []BookStock) error
	Update(ctx context.Context, book *BookStock) error
	DeleteByBookId(ctx context.Context, id string) error
	DeleteByCodes(ctx context.Context, codes []string) error
}