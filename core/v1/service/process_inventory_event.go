package service

import v1 "github.com/klovercloud-ci/core/v1"

type ProcessInventoryEvent interface {
	Listen(v1.Subject)
}