package repository

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
    "github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
)

type ChargeRepository struct {
	db *goqu.Database
}

func NewCharge(con *sql.DB) domain.ChargeRepository {
	return &ChargeRepository{
		db:goqu.New("default", con),
	}
}

func (cr *ChargeRepository) Save(ctx context.Context, charge *domain.Charge) error {
	executor := cr.db.Insert("charges").Rows(charge).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}