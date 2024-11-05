package service

import (
	"context"

	"go.opentelemetry.io/otel"
	"oss/api/oss"
)

func (s *OssService) PreSigned(ctx context.Context, req *oss.PreSignedRequest) (rp *oss.PreSignedReply, err error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "PreSigned")
	defer span.End()
	rp = &oss.PreSignedReply{}
	switch req.Category {
	case oss.PreSignedCategory_GET:
		rp.Uri, err = s.minio.PreSignedGetObject(ctx, req.Name)
	case oss.PreSignedCategory_PUT:
		rp.Uri, err = s.minio.PreSignedPutObject(ctx, req.Name)
	case oss.PreSignedCategory_HEAD:
		rp.Uri, err = s.minio.PreSignedHeadObject(ctx, req.Name)
	}
	return
}
