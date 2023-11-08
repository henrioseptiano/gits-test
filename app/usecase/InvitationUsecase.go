package usecase

import (
	"errors"
	"gits-test/models/model"
	"gits-test/models/request"
	"gits-test/models/response"

	"github.com/google/uuid"
)

func (au *AppUsecase) CreateInvitationUsecase(createInvitationRequest request.CreateInvitationRequest) (err error) {
	if createInvitationRequest.MemberID == "" {
		return errors.New("gathering id is required")
	}
	if createInvitationRequest.GatheringID == "" {
		return errors.New("gathering id is required")
	}

	member, err := au.AppRepository.GetMemberByMemberIDRepository(createInvitationRequest.MemberID)
	if err != nil {
		return errors.New("member not found")
	}

	gathering, err := au.AppRepository.GetGatheringByGatheringIDRepository(createInvitationRequest.GatheringID)
	if err != nil {
		return errors.New("gathering data not found")
	}

	newUUID, _ := uuid.NewUUID()
	invitation := model.Invitation{
		InvitationID: newUUID.String(),
		MemberID:     member.ID,
		GatheringID:  gathering.ID,
	}

	invitationID, err := au.AppRepository.CreateInvitationRepository(invitation)
	if err != nil {
		return err
	}

	invitationStatus := model.InvitationStatus{
		InvitationID: invitationID,
		Status:       1,
	}
	err = au.AppRepository.CreateInvitationStatusRepository(invitationStatus)
	if err != nil {
		return err
	}

	return nil

}

func (au *AppUsecase) AcceptNotAcceptUsecase(dataReq request.AcceptNotAcceptRequest, invitationID string) error {
	if dataReq.MemberID == "" {
		return errors.New("member id is required")
	}
	if dataReq.GatheringID == "" {
		return errors.New("gathering id is required")
	}

	member, err := au.AppRepository.GetMemberByMemberIDRepository(dataReq.MemberID)
	if err != nil {
		return errors.New("member not found")
	}

	gathering, err := au.AppRepository.GetGatheringByGatheringIDRepository(dataReq.GatheringID)
	if err != nil {
		return errors.New("gathering data not found")
	}

	invitation, err := au.AppRepository.GetInvitationByInvitationIDRepository(invitationID)
	if err != nil {
		return errors.New("invitation not found")
	}

	invitationStatus, err := au.AppRepository.GetInvitationStatusByInvitationIDRepository(invitation.ID)
	if err != nil {
		return errors.New("invitation status not found")
	}

	if invitationStatus.Status != 1 {
		var statusStr = ""
		if invitationStatus.Status == 2 {
			statusStr = "accepted"
		} else {
			statusStr = "rejected"
		}
		return errors.New("invitation status is already " + statusStr)
	}

	if dataReq.Accept {
		// if accept, create attendee
		invitationStatus.Status = 2

		err = au.AppRepository.AcceptNotAcceptRepository(invitationStatus)
		if err != nil {
			return err
		}

		newUUID, _ := uuid.NewUUID()

		attendee := model.Attendee{
			AttendeeID:  newUUID.String(),
			MemberID:    member.ID,
			GatheringID: gathering.ID,
		}
		err := au.AppRepository.CreateAttendeeRepository(attendee)
		if err != nil {
			return err
		}
	} else {
		// else, update invitation status to rejected
		invitationStatus.Status = 0

		err = au.AppRepository.AcceptNotAcceptRepository(invitationStatus)
		if err != nil {
			return err
		}
	}

	return nil
}

func (au *AppUsecase) ListInvitationUsecase() ([]response.ListInvitationResponse, error) {
	listInvitationResponse, err := au.AppRepository.ListInvitationRepository()
	if err != nil {
		return nil, err
	}

	return listInvitationResponse, nil
}
