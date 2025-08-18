package repository

import (
	"database/sql"
    "github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
)

type BookStockRepository struct {
	db *goqu.Database
}

func NewBookStock(con *sql.DB) domain.BookStockRepository {
	return &BookStockRepository{
		db:goqu.New("default", con),
	}
}

func (br *BookStockRepository) FindByBookId(ctx context.Context, id string) (result []domain.BookStock, err error) {
    dataset := br.db.From("book_stocks").Where(goqu.C("book_id").Eq(id))
    err = dataset.ScanStructsContext(ctx, &result)
    return 
}

func (br *BookStockRepository) FindByBookAndCode(ctx context.Context, id string, code string) (result domain.BookStock, err error) {
    dataset := br.db.From("book_stocks").Where(
        goqu.C("id").Eq(id), 
        goqu.C("code").Eq(code), 
    )

    _, err = dataset.ScanStructContext(ctx, &result)

    return result, err
}


func (br *BookStockRepository) Save(ctx context.Context, data []domain.BookStock) error {
    executor := br.db.Insert("book_stocks").Rows(data).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (br *BookStockRepository) Update(ctx context.Context, stock *domain.BookStock) error {
    executor := br.db.Update("book_stocks").Set(stock).Where(goqu.C("code").Eq(stock.Code)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (br *BookStockRepository) DeleteByBookId(ctx context.Context, id string) error {
	executor := br.db.Delete("book_stocks").
        Where(goqu.C("book_id").In(id)).
        Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (br *BookStockRepository) DeleteByCodes(ctx context.Context, codes []string) error {
	executor := br.db.Delete("book_stocks").
        Where(goqu.C("id").In(codes)).
        Executor()
    _, err := executor.ExecContext(ctx)
    return err
}