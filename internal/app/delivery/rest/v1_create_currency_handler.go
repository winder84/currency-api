package rest

import (
	"currency-api/internal/app/types"
	"currency-api/pkg/fibererror"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type V1CreateCurrencyRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (h Handler) v1CreateCurrencyHandler(c *fiber.Ctx) error {
	req := V1CreateCurrencyRequest{}
	err := json.Unmarshal(c.Request().Body(), &req)
	if err != nil {
		return fibererror.New(c, err)
	}

	err = h.useCase.Currency.Create(c.Context(), types.Currency{
		From: req.From,
		To:   req.To,
	})

	return fibererror.New(c, err)
}
