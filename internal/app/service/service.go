package service

import (
	"context"
	"currency-api/internal/app/config"
	"currency-api/internal/app/repository"
	"currency-api/internal/app/types"
)

type UseCase struct {
	Currency Currencier
}

type Currencier interface {
	Create(ctx context.Context, currency types.Currency) error
	Update(ctx context.Context, currency types.Currency) error
	Get(ctx context.Context, currency types.Currency) (types.Currency, error)
	GetAll(ctx context.Context) ([]types.Currency, error)
	UpdateAll(ctx context.Context) error
}

func New(storage *repository.Storage, apiConfig config.Api) *UseCase {
	return &UseCase{
		Currency: newCurrency(storage, apiConfig),
	}
}
