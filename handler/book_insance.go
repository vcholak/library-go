package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CopiesTotal returns total number of BookInstances
func (h *Handler) CopiesTotal(c echo.Context) error {

  copies, err := h.copyStore.BookInstanceCount()
  if err != nil {
    panic("Failed to fetch book instance count")
  }

	c.Response().Header().Set("X-Result-Count", strconv.FormatInt(copies, 10))
  c.Response().Header().Set("Access-Control-Expose-Headers", "X-Result-Count")

	return c.NoContent(http.StatusOK)
}

// AvailableCopiesTotal returns total number of BookInstances
func (h *Handler) AvailableCopiesTotal(c echo.Context) error {

  available_copies, err := h.copyStore.AvailableBookInstanceCount()
  if err != nil {
    panic("Failed to fetch book instance count")
  }

  c.Response().Header().Set("X-Result-Count", strconv.FormatInt(available_copies, 10))
  c.Response().Header().Set("Access-Control-Expose-Headers", "X-Result-Count")

	return c.NoContent(http.StatusOK)
}

// Copies returns all book copies
func (h *Handler) Copies(c echo.Context) error {
	//
	return nil
}
