package repository

import (
	"gits-test/models/model"
	"gits-test/models/response"
)

func (ar *AppRepository) CreateGatheringRepository(gatheringModel model.Gathering) (int, error) {
	if err := ar.DB.Create(&gatheringModel).Error; err != nil {
		return 0, err
	}
	return gatheringModel.ID, nil
}

func (ar *AppRepository) CreateGatheringTypeRepository(gatheringTypeModel model.GatheringType) error {
	if err := ar.DB.Create(&gatheringTypeModel).Error; err != nil {
		return err
	}
	return nil
}

func (ar *AppRepository) GetGatheringByGatheringIDRepository(gatheringID string) (model.Gathering, error) {
	var gatheringModel model.Gathering
	if err := ar.DB.Where("gathering_id = ?", gatheringID).First(&gatheringModel).Error; err != nil {
		return gatheringModel, err
	}
	return gatheringModel, nil
}

func (ar *AppRepository) ListGatheringRepository() ([]response.ListGatheringResponse, error) {
	listGathering := make([]response.ListGatheringResponse, 0)
	query := ar.DB.Table("gatherings as g").Joins("JOIN gathering_types as gt ON g.id = gt.gathering_id")
	query.Where("g.status = ?", 1)
	query.Select("g.gathering_id, name as gathering_name, creator, DATE_FORMAT(scheduled_at, '%Y-%m-%d %H:%i:%s') AS scheduled_at, location, gt.type as gathering_type")
	err := query.Find(&listGathering).Error
	if err != nil {
		return listGathering, err
	}
	return listGathering, nil
}
