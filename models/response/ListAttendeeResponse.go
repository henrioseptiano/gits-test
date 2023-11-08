package response

type ListAttendeeResponse struct {
	AttendeeID  string `json:"attendee_id"`
	MemberID    string `json:"member_id"`
	GatheringID string `json:"gathering_id"`
}
