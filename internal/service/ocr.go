package service

import (
	"context"

	"github.com/go-cinch/common/copierx"
	"go.opentelemetry.io/otel"
	"oss/api/oss"
)

func (s *OssService) Ocr(ctx context.Context, req *oss.OcrRequest) (rp *oss.OcrReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "Ocr")
	defer span.End()
	rp = &oss.OcrReply{}
	list := s.ocr.Ocr(ctx, req.List...)
	copierx.Copy(&rp.List, list)
	return
}