package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CopiesTotal returns total number of BookInstances
func (h *Handler) CopiesTotal(c echo.Context) error {

  copies := 10

	c.Response().Header().Set("X-Result-Count", strconv.Itoa(copies))
  c.Response().Header().Set("Access-Control-Expose-Headers", "X-Result-Count")

	return c.NoContent(http.StatusOK)
}

// AvailableCopiesTotal returns total number of BookInstances
func (h *Handler) AvailableCopiesTotal(c echo.Context) error {

  available_copies := 5

  c.Response().Header().Set("X-Result-Count", strconv.Itoa(available_copies))
  c.Response().Header().Set("Access-Control-Expose-Headers", "X-Result-Count")

	return c.NoContent(http.StatusOK)
}

// Copies returns all book copies
func (h *Handler) Copies(c echo.Context) error {
	//
	return nil
}
