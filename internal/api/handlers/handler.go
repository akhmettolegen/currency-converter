package handlers

import "github.com/akhmettolegen/currency-converter/pkg/application"

type Handler struct {
	App *application.Application
}

func Get(app *application.Application) *Handler{

	return &Handler{
		App: app,
	}
}
