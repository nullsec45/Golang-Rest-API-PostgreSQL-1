package repository

import (
	"database/sql"
    "github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
    "time"
)

type BookRepository struct {
	db *goqu.Database
}

func NewBook(con *sql.DB) domain.BookRepository {
	return &BookRepository{
		db:goqu.New("default", con),
	}
}

func (br *BookRepository) FindAll(ctx context.Context) (result []domain.Book, err error) {
    dataset := br.db.From("books").Where(goqu.C("deleted_at").IsNull())
    err = dataset.ScanStructsContext(ctx, &result)
    return 
}

func (br *BookRepository) FindById(ctx context.Context, id string) (result domain.Book, err error) {
    dataset := br.db.From("books").Where(
        goqu.C("id").Eq(id), 
        goqu.C("deleted_at").IsNull(),
    )
    found, err := dataset.ScanStructContext(ctx, &result)
    if !found {
        return result, sql.ErrNoRows
    }
    return result, err
}


func (br *BookRepository) Save(ctx context.Context, b *domain.Book) error {
    executor := br.db.Insert("books").Rows(b).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (br *BookRepository) Update(ctx context.Context, b *domain.Book) error {
    executor := br.db.Update("books").Set(b).Where(goqu.C("id").Eq(b.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (br *BookRepository) Delete(ctx context.Context, id string) error {
    executor := br.db.Update("books").
        Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).
        Where(goqu.C("id").Eq(id)).
        Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

