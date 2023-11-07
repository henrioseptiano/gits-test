package repository

import "gits-test/models/model"

func (ar *AppRepository) CreateMemberRepository(member model.Member) error {
	if err := ar.DB.Create(&member).Error; err != nil {
		return err
	}
	return nil
}
