package ocr

import "context"

type API interface {
	OcrPredict(ctx context.Context, condition *Req) (*Resp, error)
	HelmetPredict(ctx context.Context, condition *Req) (*Resp, error)
}
