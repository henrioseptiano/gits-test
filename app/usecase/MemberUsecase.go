package usecase

import (
	"errors"
	"gits-test/models/model"
	"gits-test/models/request"
	"gits-test/models/response"
	"gits-test/utils"

	"github.com/google/uuid"
)

func (au *AppUsecase) CreateMemberUsecase(createMemberRequest request.CreateMemberRequest) error {
	if createMemberRequest.FirstName == "" {
		return errors.New("First Name is required")
	}
	if createMemberRequest.LastName == "" {
		return errors.New("Last Name is required")
	}
	if createMemberRequest.Email == "" {
		return errors.New("Email is required")
	}
	if createMemberRequest.Address == "" {
		return errors.New("Address is required")
	}
	if createMemberRequest.PhoneNumber == "" {
		return errors.New("Phone Number is required")
	}
	if err := utils.ValidateEmail(createMemberRequest.Email); err != nil {
		return err
	}

	if !utils.IsValidPhoneNumber(createMemberRequest.PhoneNumber) {
		return errors.New("Invalid Phone Number format")
	}

	newMemberId := uuid.New().String()
	member := model.Member{
		MemberID:    newMemberId,
		FirstName:   createMemberRequest.FirstName,
		LastName:    createMemberRequest.LastName,
		Email:       createMemberRequest.Email,
		Address:     createMemberRequest.Address,
		PhoneNumber: createMemberRequest.PhoneNumber,
	}
	if err := au.AppRepository.CreateMemberRepository(member); err != nil {
		return err
	}

	return nil
}

func (au *AppUsecase) ListMemberUsecase() ([]response.ListMemberResponse, error) {
	listMemberResponse, err := au.AppRepository.ListMemberRepository()
	if err != nil {
		return nil, err
	}

	return listMemberResponse, nil
}
