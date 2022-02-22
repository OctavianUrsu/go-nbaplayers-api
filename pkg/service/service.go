package service

import (
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/helpers"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/store"
)

type Service struct {
	helpers *helpers.Helpers
	store   *store.Store
}

// Constructor for dependency injection
func NewService(h *helpers.Helpers, r *store.Store) *Service {
	return &Service{h, r}
}
