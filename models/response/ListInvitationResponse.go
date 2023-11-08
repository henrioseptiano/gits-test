package response

type ListInvitationResponse struct {
	InvitationID     string `json:"invitation_id"`
	GatheringID      string `json:"gathering_id"`
	MemberID         string `json:"member_id"`
	InvitationStatus string `json:"invitation_status"`
}
