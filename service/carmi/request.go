package carmi

type CarmiGenRequest struct {
	Type     string `json:"type"`
	EndTime  string `json:"end_time"`
	Quantity int    `json:"quantity"`
	Reward   int32  `json:"reward"`
	Level    int32  `json:"level"`
}
