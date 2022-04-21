package rest

import (
	"currency-api/internal/app/types"
	"currency-api/pkg/fibererror"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type V1ConvertCurrencyRequest struct {
	From  string  `json:"from"`
	To    string  `json:"to"`
	Value float64 `json:"value"`
}

type V1ConvertCurrencyResponse struct {
	Result float64 `json:"result"`
}

func (h Handler) v1ConvertCurrencyHandler(c *fiber.Ctx) error {
	req := V1ConvertCurrencyRequest{}
	err := json.Unmarshal(c.Request().Body(), &req)
	if err != nil {
		return fibererror.New(c, err)
	}

	storageCurr, err := h.useCase.Currency.Get(c.Context(), types.Currency{
		From: req.From,
		To:   req.To,
	})
	if err != nil {
		return fibererror.New(c, err)
	}

	resp := &V1ConvertCurrencyResponse{
		Result: req.Value * storageCurr.Well,
	}

	return c.JSON(resp)
}
