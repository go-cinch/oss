package ocr

import "context"

type API interface {
	Ocr(ctx context.Context, image string) ([][]Resp, error)
}
