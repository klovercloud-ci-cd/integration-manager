package api

import "github.com/labstack/echo/v4"

// Pipeline pipeline api operations
type Pipeline interface {
	Get(context echo.Context) error
	Update(context echo.Context) error
	Create(context echo.Context) error
}
