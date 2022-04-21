package service

import (
	"context"
	"currency-api/pkg/open_exchange_client"

	"currency-api/internal/app/config"
	"currency-api/internal/app/repository"
	"currency-api/internal/app/types"
)

type currency struct {
	CSConfig config.Api
	storage  *repository.Storage
}

func newCurrency(storage *repository.Storage, CSConfig config.Api) *currency {
	return &currency{
		storage:  storage,
		CSConfig: CSConfig,
	}
}

func (c *currency) Get(ctx context.Context, currency types.Currency) (types.Currency, error) {
	return c.storage.Currency.GetCurrency(ctx, currency)
}

func (c *currency) Create(ctx context.Context, currency types.Currency) error {
	return c.storage.Currency.CreateCurrency(ctx, currency)
}

func (c *currency) Update(ctx context.Context, currency types.Currency) error {
	return c.storage.Currency.UpdateCurrency(ctx, currency)
}

func (c *currency) GetAll(ctx context.Context) ([]types.Currency, error) {
	return c.storage.Currency.GetCurrencies(ctx)
}

func (c *currency) UpdateAll(ctx context.Context) error {
	oeClient := open_exchange_client.New(open_exchange_client.Config{
		Url:   c.CSConfig.Url,
		AppId: c.CSConfig.AppId,
	})

	storageCurrencies, err := c.GetAll(ctx)
	if err != nil {
		return err
	}

	oeCurrencies := make([]open_exchange_client.Currency, 0, len(storageCurrencies))
	for _, sc := range storageCurrencies {
		oeCurrencies = append(oeCurrencies, open_exchange_client.Currency{
			From: sc.From,
			To:   sc.To,
		})
	}

	resp, err := oeClient.GetRates(oeCurrencies)
	if err != nil {
		return err
	}

	for _, d := range resp.Data {
		err := c.Update(ctx, types.Currency{
			From: d.From,
			To:   d.To,
			Well: d.Well,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
