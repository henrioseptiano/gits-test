package usecase

import "gits-test/models/response"

func (au *AppUsecase) ListAttendeeUsecase() (listAttendeeResponse []response.ListAttendeeResponse, err error) {
	attendee, err := au.AppRepository.ListAttendeeRepository()
	if err != nil {
		return nil, err
	}

	return attendee, nil
}
