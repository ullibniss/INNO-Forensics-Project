package api

import (
	"errors"
	"mac_identifier/internal/applications/mac"

	"github.com/gofiber/fiber/v2"
)

type MacHandler struct {
	MacService mac.Service
}

func NewMacHandler(macService mac.Service) *MacHandler {
	return &MacHandler{MacService: macService}
}

func (h *MacHandler) addMacRouter(api fiber.Router) {
	api.Get("/api/lookup", h.lookup)
}

func (h *MacHandler) lookup(c *fiber.Ctx) error {
	macAddr := c.Query("mac")
	if macAddr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "mac query param required"})
	}

	result, err := h.MacService.Lookup(macAddr)
	if err != nil {
		if errors.Is(err, mac.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "vendor not found"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"mac": result.MAC, "vendor": result.Vendor})
}
