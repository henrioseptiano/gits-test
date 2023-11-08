package repository

import (
	"gits-test/models/model"
	"gits-test/models/response"
)

func (ar *AppRepository) CreateMemberRepository(member model.Member) error {
	if err := ar.DB.Create(&member).Error; err != nil {
		return err
	}
	return nil
}

func (ar *AppRepository) GetMemberByMemberIDRepository(memberID string) (model.Member, error) {
	var member model.Member
	if err := ar.DB.Where("member_id = ?", memberID).First(&member).Error; err != nil {
		return member, err
	}
	return member, nil
}

func (ar *AppRepository) ListMemberRepository() ([]response.ListMemberResponse, error) {
	listMember := make([]response.ListMemberResponse, 0)
	query := ar.DB.Table("members").Select("member_id, CONCAT(first_name, ' ', last_name) as member_name, email as member_email, address as member_address, phone_number as member_phone")
	err := query.Find(&listMember).Error
	if err != nil {
		return listMember, err
	}
	return listMember, nil
}
