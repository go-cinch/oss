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

func (Ocr) OcrPredict(_ context.Context, _ *ocr.Req) (*ocr.Resp, error) {
	time.Sleep(100 * time.Millisecond)
	return &ocr.Resp{
		List: []ocr.Detail{
			{
				Original: "1.jpg",
				Points: []ocr.Point{
					{
						Confidence: "0",
						Text:       "1000",
						TextRegion: "[[0,0]]",
					},
				},
			},
		},
	}, nil
}

func (Ocr) HelmetPredict(_ context.Context, _ *ocr.Req) (*ocr.Resp, error) {
	time.Sleep(100 * time.Millisecond)
	return &ocr.Resp{
		List: []ocr.Detail{
			{
				Original: "2.jpg",
				Points: []ocr.Point{
					{
						Confidence: "0",
						Text:       "1000",
						TextRegion: "[[0,0]]",
					},
				},
			},
		},
	}, nil
}
