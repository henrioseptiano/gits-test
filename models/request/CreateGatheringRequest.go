package request

type CreateGatheringRequest struct {
	Creator     string `json:"creator"`
	Type        string `json:"type"`
	ScheduledAt string `json:"scheduled_at"`
	Name        string `json:"name"`
	Location    string `json:"location"`
}
