package response

type ListGatheringResponse struct {
	GatheringID   string `json:"gathering_id"`
	GatheringName string `json:"gathering_name"`
	Creator       string `json:"creator"`
	ScheduledAt   string `json:"scheduled_at"`
	Location      string `json:"location"`
	GatheringType string `json:"gathering_type"`
}
