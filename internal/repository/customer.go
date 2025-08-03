package repository

import (
    "context"
    "database/sql"
    "time"
    "github.com/doug-martin/goqu/v9"
    "github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
)

type CustomerRepository struct {
    db *goqu.Database
}

func NewCustomer(con *sql.DB) domain.CustomerRepository {
    return &CustomerRepository{
        db: goqu.New("postgres", con),
    }
}

func (cr *CustomerRepository) FindAll(ctx context.Context) (result []domain.Customer, err error) {
    dataset := cr.db.From("customers").Where(goqu.C("deleted_at").IsNull())
    err = dataset.ScanStructsContext(ctx, &result)
    return 
}

func (cr *CustomerRepository) FindById(ctx context.Context, id string) (result domain.Customer, err error) {
    dataset := cr.db.From("customers").Where(
        goqu.C("id").Eq(id), 
        goqu.C("deleted_at").IsNull(),
    )
    found, err := dataset.ScanStructContext(ctx, &result)
    if !found {
        return result, sql.ErrNoRows
    }
    return result, err
}


func (cr *CustomerRepository) Save(ctx context.Context, c *domain.Customer) error {
    executor := cr.db.Insert("customers").Rows(c).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (cr *CustomerRepository) Update(ctx context.Context, c *domain.Customer) error {
    executor := cr.db.Update("customers").Set(c).Where(goqu.C("id").Eq(c.ID)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (cr *CustomerRepository) Delete(ctx context.Context, id string) error {
    executor := cr.db.Update("customers").
        Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).
        Where(goqu.C("id").Eq(id)).
        Executor()
    _, err := executor.ExecContext(ctx)
    return err
}
