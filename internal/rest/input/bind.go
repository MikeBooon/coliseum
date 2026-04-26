package input

import "github.com/labstack/echo/v5"

func BindAndValidate[T any](c *echo.Context, body *T) error {
	if err := c.Bind(body); err != nil {
		return err
	}
	return c.Validate(body)
}
