package domain

import (
	"context"
	"database/sql"
)

type Charge struct {
	Id string `db:"id"`
	JournalId string `db:"journal_id"`
	DaysLate int `db:"days_late"`
	DailyLateFee int `db:"daily_late_fee"`
	UserId string `db:"user_id"`
	Total int `db:"total"`
	CreatedAt sql.NullTime `db:"created_at"`
}

type ChargeRepository interface {
	Save(ctx context.Context, charge *Charge) error
}