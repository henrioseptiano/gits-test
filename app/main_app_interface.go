package app

import (
	"gits-test/models/model"
	"gits-test/models/request"

	"github.com/gin-gonic/gin"
)

type AppHandlerInterface interface {
	CreateMemberHandler(c *gin.Context)
	CreateGatheringHandler(c *gin.Context)
}

type AppUsecaseInterface interface {
	CreateMemberUsecase(request.CreateMemberRequest) error
	CreateGatheringUsecase(request.CreateGatheringRequest) error
}

type AppRepositoryInterface interface {
	CreateMemberRepository(model.Member) error
	CreateGatheringRepository(model.Gathering) (int, error)
	CreateGatheringTypeRepository(model.GatheringType) error
}
