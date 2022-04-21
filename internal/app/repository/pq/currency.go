package pq

import (
	"context"
	"currency-api/internal/app/types"
)

func (rep *PostgresRepository) GetCurrency(ctx context.Context, currency types.Currency) (respCurrency types.Currency, err error) {
	query := `
        select currency_from, currency_to, well, updated_at from currencies
        where currency_from = $1 and currency_to = $2`

	return respCurrency, rep.db.GetContext(ctx, &respCurrency, query, currency.From, currency.To)
}

func (rep *PostgresRepository) GetCurrencies(ctx context.Context) (currencies []types.Currency, err error) {
	query := `select currency_from, currency_to, well, updated_at from currencies`

	return currencies, rep.db.SelectContext(ctx, &currencies, query)
}

func (rep *PostgresRepository) CreateCurrency(ctx context.Context, currency types.Currency) (err error) {
	query := `
		insert into currencies(currency_from, currency_to)
		values (:currency_from, :currency_to)`

	_, err = rep.db.NamedExecContext(ctx, query, currency)
	return err
}

func (rep *PostgresRepository) UpdateCurrency(ctx context.Context, currency types.Currency) (err error) {
	query := `
		update currencies set well=:well,updated_at=now()
		where currency_from=:currency_from and currency_to=:currency_to`

	_, err = rep.db.NamedExecContext(ctx, query, currency)
	return err
}
