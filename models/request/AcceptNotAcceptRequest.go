package request

type AcceptNotAcceptRequest struct {
	Accept      bool   `json:"accept"`
	MemberID    string `json:"member_id"`
	GatheringID string `json:"gathering_id"`
}
