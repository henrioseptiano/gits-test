package handler

import "gits-test/app"

type AppHandler struct {
	AppUsecase app.AppUsecaseInterface
}

func NewAppHandler(appUsecase app.AppUsecaseInterface) *AppHandler {
	return &AppHandler{
		AppUsecase: appUsecase,
	}
}
