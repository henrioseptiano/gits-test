package repository

import (
	"gits-test/models/model"
	"gits-test/models/response"
)

func (ar *AppRepository) CreateInvitationRepository(invitation model.Invitation) (int, error) {
	err := ar.DB.Create(&invitation).Error
	if err != nil {
		return 0, err
	}
	return invitation.ID, nil
}

func (ar *AppRepository) CreateInvitationStatusRepository(invitationStatus model.InvitationStatus) error {
	err := ar.DB.Create(&invitationStatus).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *AppRepository) AcceptNotAcceptRepository(invitationStatus model.InvitationStatus) error {
	err := ar.DB.Save(&invitationStatus).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *AppRepository) GetInvitationByInvitationIDRepository(invitationID string) (model.Invitation, error) {
	var invitation model.Invitation
	err := ar.DB.Where("invitation_id = ?", invitationID).First(&invitation).Error
	if err != nil {
		return invitation, err
	}
	return invitation, nil
}

func (ar *AppRepository) GetInvitationStatusByInvitationIDRepository(invitationID int) (model.InvitationStatus, error) {
	var invitationStatus model.InvitationStatus
	err := ar.DB.Where("invitation_id = ?", invitationID).First(&invitationStatus).Error
	if err != nil {
		return invitationStatus, err
	}
	return invitationStatus, nil
}

func (ar *AppRepository) ListInvitationRepository() ([]response.ListInvitationResponse, error) {
	listInvitation := make([]response.ListInvitationResponse, 0)
	query := ar.DB.Table("invitations as i").Joins("JOIN invitation_statuses as is1 ON i.id = is1.invitation_id")
	query.Joins("JOIN gatherings as g ON i.gathering_id = g.id")
	query.Joins("JOIN members as u ON i.member_id = u.id")
	query.Where("g.status = ?", 1)
	query.Select("i.invitation_id, u.member_id, g.gathering_id, CASE WHEN is1.status = 1 THEN 'Created' WHEN is1.status = 2 THEN 'Accepted' ELSE 'Rejected' END AS invitation_status")
	err := query.Find(&listInvitation).Error
	if err != nil {
		return listInvitation, err
	}
	return listInvitation, nil
}
