package health

import (
	"context"
	"net/http"
	"time"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/health"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	database *persistence.Database
}

func NewHandler(database *persistence.Database) *Handler {
	return &Handler{
		database: database,
	}
}

func (h *Handler) Handle(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 30*time.Second)
	defer cancel()

	db, err := h.database.Instance.DB()
	if err != nil {
		status := health.New(err)
		return c.JSON(http.StatusInternalServerError, status)
	}

	if err := db.PingContext(ctx); err != nil {
		status := health.New(err)
		return c.JSON(http.StatusInternalServerError, status)
	}

	return c.JSON(http.StatusOK, health.New(nil))
}
