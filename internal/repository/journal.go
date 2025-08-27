package repository

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
    "github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
)

type JournalRepository struct {
	db *goqu.Database
}

func NewJournal(con *sql.DB) domain.JournalRepository {
	return &JournalRepository{
		db:goqu.New("default", con),
	}
}

func (j JournalRepository) Find(ctx context.Context, se domain.JournalSearch) (result []domain.Journal, err error){
	dataset := j.db.From("journals")

	if se.CustomerId != "" {
		dataset = dataset.Where(goqu.C("customer_id").Eq(se.CustomerId))
	}

	if se.Status != "" {
		dataset = dataset.Where(goqu.C("status").Eq(se.Status))
	}

	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (j JournalRepository) FindById(ctx context.Context, id string) (result domain.Journal, err error) {
	dataset := j.db.From("journals").Where(goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (j JournalRepository) Save (ctx context.Context, journal *domain.Journal) error {
	executor := j.db.Insert("journals").Rows(journal).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (j JournalRepository) Update (ctx context.Context, journal *domain.Journal) error {
	 executor := j.db.Update("journals").Set(journal).Where(goqu.C("id").Eq(journal.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}