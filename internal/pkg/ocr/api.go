package ocr

import "context"

type API interface {
	PredictAll(ctx context.Context, image string) (*Resp, error)
}
