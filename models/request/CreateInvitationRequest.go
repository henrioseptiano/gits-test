package request

type CreateInvitationRequest struct {
	MemberID    string `json:"member_id"`
	GatheringID string `json:"gathering_id"`
}
