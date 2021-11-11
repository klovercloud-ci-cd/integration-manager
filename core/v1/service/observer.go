package service

import "github.com/klovercloud-ci/core/v1"

// Observer Observer related operations.
type Observer interface {
	Listen(subject v1.Subject)
}
