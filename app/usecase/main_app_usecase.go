package usecase

import "gits-test/app"

type AppUsecase struct {
	AppRepository app.AppRepositoryInterface
}

func NewAppUsecase(appRepository app.AppRepositoryInterface) *AppUsecase {
	return &AppUsecase{
		AppRepository: appRepository,
	}
}
