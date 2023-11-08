package repository

import (
	"gits-test/models/model"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateMemberRepository(t *testing.T) {

	maem.Mock.On("Create", mock.Anything).Return(&gorm.DB{})

	memberID, _ := uuid.NewUUID()
	member := model.Member{
		MemberID:    memberID.String(),
		FirstName:   "budi",
		LastName:    "santoso",
		Email:       "budi.santoso@gmail.com",
		Address:     "jakarta",
		PhoneNumber: "08123456789",
	}
	err := mainAppEntity.CreateMemberRepository(member)

	maem.Mock.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}
}

func TestGetMemberByMemberIDRepository(t *testing.T) {
	// Initialize the mock
	maem.Mock.On("Where", "member_id = 4ec1b06b-7ddc-11ee-9739-3822e2435ffd").Return(&gorm.DB{})
	maem.Mock.On("First", mock.Anything).Return(&gorm.DB{})
	_, err := mainAppEntity.GetMemberByMemberIDRepository("4ec1b06b-7ddc-11ee-9739-3822e2435ffd")

	maem.Mock.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

}

func TestListMemberRepository(t *testing.T) {
	// Initialize the mock
	maem.Mock.On("Table", "members").Return(&gorm.DB{})
	maem.Mock.On("Select", "member_id, CONCAT(first_name, ' ', last_name) as member_name, email as member_email, address as member_address, phone_number as member_phone").Return(&gorm.DB{})
	maem.Mock.On("Find", mock.Anything).Return(&gorm.DB{})

	_, err := mainAppEntity.ListMemberRepository()

	maem.Mock.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

}
