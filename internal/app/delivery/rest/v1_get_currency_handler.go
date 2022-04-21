package rest

import (
	"currency-api/pkg/fibererror"
	"github.com/gofiber/fiber/v2"
)

type V1CurrencyResponse struct {
	Currencies []Currency `json:"pairs"`
}

type Currency struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (h Handler) v1GetCurrencyHandler(c *fiber.Ctx) error {
	currencies, err := h.useCase.Currency.GetAll(c.Context())
	if err != nil {
		return fibererror.New(c, err)
	}

	resp := V1CurrencyResponse{
		Currencies: make([]Currency, 0, len(currencies)),
	}

	for _, cc := range currencies {
		resp.Currencies = append(resp.Currencies, Currency{
			From: cc.From,
			To:   cc.To,
		})
	}

	return c.JSON(resp)
}
