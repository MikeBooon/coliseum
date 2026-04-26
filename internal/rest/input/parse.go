package input

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func ParseUUID(in string) (uuid.UUID, error) {
	id, err := uuid.Parse(in)
	if err != nil {
		return uuid.UUID{}, echo.NewHTTPError(http.StatusBadRequest, "invalid uuid")
	}
	return id, nil
}
