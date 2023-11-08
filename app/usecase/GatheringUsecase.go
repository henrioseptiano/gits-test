package usecase

import (
	"errors"
	"gits-test/models/model"
	"gits-test/models/request"
	"gits-test/models/response"
	"time"

	"github.com/google/uuid"
)

func (au *AppUsecase) CreateGatheringUsecase(createGatheringRequest request.CreateGatheringRequest) error {
	if createGatheringRequest.Creator == "" {
		return errors.New("Creator is required")
	}

	if createGatheringRequest.Type == "" {
		return errors.New("Type is required")
	}

	if createGatheringRequest.ScheduledAt == "" {
		return errors.New("ScheduledAt is required")
	}

	if createGatheringRequest.Name == "" {
		return errors.New("Name is required")
	}

	if createGatheringRequest.Location == "" {
		return errors.New("Location is required")
	}

	scheduledAtTime, err := time.Parse("2006-01-02 15:04:05", createGatheringRequest.ScheduledAt)
	if err != nil {
		return errors.New("scheduled_at format must be yyyy-MM-dd hh:mm:ss")
	}
	scheduledAtTimeNano := scheduledAtTime.UnixNano()
	now := time.Now().UnixNano()

	if scheduledAtTimeNano < now {
		return errors.New("scheduled_at must be greater than now")
	}
	uuidNew, _ := uuid.NewUUID()

	gatheringModel := model.Gathering{
		GatheringID: uuidNew.String(),
		Creator:     createGatheringRequest.Creator,
		Name:        createGatheringRequest.Name,
		Location:    createGatheringRequest.Location,
		Status:      true,
		ScheduledAt: scheduledAtTime,
	}
	gatheringId, err := au.AppRepository.CreateGatheringRepository(gatheringModel)
	if err != nil {
		return err
	}

	gatheringTypeModel := model.GatheringType{
		GatheringID: gatheringId,
		Type:        createGatheringRequest.Type,
	}

	if err := au.AppRepository.CreateGatheringTypeRepository(gatheringTypeModel); err != nil {
		return err
	}

	return nil
}

func (au *AppUsecase) ListGatheringUsecase() ([]response.ListGatheringResponse, error) {
	listGathering, err := au.AppRepository.ListGatheringRepository()
	if err != nil {
		return listGathering, err
	}

	return listGathering, nil
}
