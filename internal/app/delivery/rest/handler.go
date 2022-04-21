package rest

import (
	"currency-api/internal/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type Handler struct {
	useCase  *service.UseCase
	fiberAPI *fiber.App
}

func New(useCase *service.UseCase) *Handler {
	h := &Handler{
		useCase: useCase,
	}

	api := fiber.New()
	api.Get("/api/v1/currency", h.v1GetCurrencyHandler)
	api.Put("/api/v1/currency", h.v1ConvertCurrencyHandler)
	api.Post("/api/v1/currency", h.v1CreateCurrencyHandler)

	h.fiberAPI = api
	return h
}

func (h *Handler) Listen(port string) error {
	return errors.Wrap(h.fiberAPI.Listen(port), "listen handler api")
}
