package app

import (
	"gits-test/models/model"
	"gits-test/models/request"
	"gits-test/models/response"

	"github.com/gin-gonic/gin"
)

type AppHandlerInterface interface {
	//Create
	CreateMemberHandler(c *gin.Context)
	CreateGatheringHandler(c *gin.Context)
	CreateInvitationHandler(c *gin.Context)
	//Update
	AcceptNotAcceptHandler(c *gin.Context)
	//List
	ListMemberHandler(c *gin.Context)
	ListGatheringHandler(c *gin.Context)
	ListInvitationHandler(c *gin.Context)
	ListAttendeeHandler(c *gin.Context)
}

type AppUsecaseInterface interface {
	//Create
	CreateMemberUsecase(request.CreateMemberRequest) error
	CreateGatheringUsecase(request.CreateGatheringRequest) error
	CreateInvitationUsecase(request.CreateInvitationRequest) error
	//Update
	AcceptNotAcceptUsecase(request.AcceptNotAcceptRequest, string) error
	//List
	ListMemberUsecase() ([]response.ListMemberResponse, error)
	ListGatheringUsecase() ([]response.ListGatheringResponse, error)
	ListInvitationUsecase() ([]response.ListInvitationResponse, error)
	ListAttendeeUsecase() ([]response.ListAttendeeResponse, error)
}

type AppRepositoryInterface interface {
	//Create
	CreateMemberRepository(model.Member) error
	CreateGatheringRepository(model.Gathering) (int, error)
	CreateGatheringTypeRepository(model.GatheringType) error
	CreateInvitationRepository(model.Invitation) (int, error)
	CreateInvitationStatusRepository(model.InvitationStatus) error
	CreateAttendeeRepository(model.Attendee) error

	//Get
	GetMemberByMemberIDRepository(string) (model.Member, error)
	GetGatheringByGatheringIDRepository(string) (model.Gathering, error)
	GetInvitationByInvitationIDRepository(string) (model.Invitation, error)
	GetInvitationStatusByInvitationIDRepository(int) (model.InvitationStatus, error)

	//List
	ListMemberRepository() ([]response.ListMemberResponse, error)
	ListGatheringRepository() ([]response.ListGatheringResponse, error)
	ListInvitationRepository() ([]response.ListInvitationResponse, error)
	ListAttendeeRepository() ([]response.ListAttendeeResponse, error)

	//Update
	AcceptNotAcceptRepository(model.InvitationStatus) error
}
