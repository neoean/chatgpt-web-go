package draw

const (
	DrawTypeOpenAI = "openai"
	DrawTypeSD     = "sd"
)

type Request struct {
	Prompt   string `json:"prompt"`
	Quantity int64  `json:"quantity"`
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
	Quality  int64  `json:"quality"`
	Steps    int64  `json:"steps"`
	Style    string `json:"style"`
	Image    string `json:"image"`
	DrawType string `json:"draw_type"`
}

type Response struct {
}
