package repository

import (
	"context"
	"currency-api/internal/app/repository/pq"
	"currency-api/internal/app/types"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	Currency Currency
}

type Currency interface {
	GetCurrency(ctx context.Context, currency types.Currency) (types.Currency, error)
	GetCurrencies(ctx context.Context) ([]types.Currency, error)
	CreateCurrency(ctx context.Context, currency types.Currency) error
	UpdateCurrency(ctx context.Context, currency types.Currency) error
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		Currency: pq.New(db),
	}
}
