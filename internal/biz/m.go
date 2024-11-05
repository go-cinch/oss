package biz

type OcrItem struct {
	Boxes        string `json:"boxes"`
	Confidence   string `json:"confidence"`
	Text         string `json:"text"`
	ParseLatency int64  `json:"parseLatency"`
	OcrLatency   int64  `json:"ocrLatency"`
	Msg          string `json:"msg"`
	Base64Data   string `json:"-"`
}
