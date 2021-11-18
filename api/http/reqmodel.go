package http

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	idPathParam = "id"

	startDateQParam = "start"
	endDateQParam   = "end"

	invalid = "is invalid"
)

func getPathParamAsInt(ctx echo.Context, pParam string) (int, error) {
	text := ctx.Param(pParam)
	parsed, err := strconv.Atoi(text)
	if err != nil {
		return 0, validateError{pParam + " [" + text + "]": "has invalid format"}
	}

	return parsed, nil
}

func parseDate(ctx echo.Context, qParam string) (time.Time, error) {
	text := ctx.QueryParam(qParam)
	if text == "" {
		return time.Time{}, validateError{qParam: "date " + invalid}
	}

	const layout = "2006-01-02"
	t, err := time.Parse(layout, text)
	if err != nil {
		return time.Time{}, validateError{qParam: "date must be in format '" + layout + "'"}
	}

	return t, nil
}

func validateStartEndDate(s, e time.Time) error {
	if e.Before(s) {
		return validateError{startDateQParam: "date must be before " + endDateQParam + " date"}
	}

	return nil
}
