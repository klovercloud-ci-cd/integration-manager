package api

import "github.com/labstack/echo/v4"

// Git git api operations
type Git interface {
	ListenEvent(context echo.Context) error
}