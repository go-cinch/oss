package biz

type OcrResult struct {
	Original     string    `json:"original"`
	ParseLatency int64     `json:"parseLatency"`
	OcrLatency   int64     `json:"ocrLatency"`
	Msg          string    `json:"msg"`
	List         []OcrItem `json:"list"`
}

type OcrItem struct {
	Boxes      string `json:"boxes"`
	Confidence string `json:"confidence"`
	Text       string `json:"text"`
}
