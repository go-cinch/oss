package mock

import (
	"context"
	"time"

	"oss/internal/pkg/ocr"
)

type Ocr struct {
}

func NewOcr() (api ocr.API, err error) {
	return &Ocr{}, nil
}

func (Ocr) Ocr(_ context.Context, _ string) (rp *ocr.Resp, err error) {
	time.Sleep(100 * time.Millisecond)
	return &ocr.Resp{
		Confidence: 0,
		Text:       "1000",
		TextRegion: [][]int{{0, 0}},
	}, nil
}
