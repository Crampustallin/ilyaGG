package handlers

import "github.com/Crampustallin/ilyaGG/internals/services"

type Handler struct {
	service *services.Service
}

func New() *Handler {
	return &Handler{
		service: services.New(),
	}
}
