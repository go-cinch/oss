package service

import (
	"github.com/go-cinch/common/minio"
	"github.com/go-cinch/common/worker"
	"github.com/google/wire"
	"oss/api/oss"
	"oss/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewOssService)

// OssService is a oss service.
type OssService struct {
	oss.UnimplementedOssServer

	task  *worker.Worker
	minio *minio.Minio
	oss   *biz.OssUseCase
}

// NewOssService new a service.
func NewOssService(task *worker.Worker, minio *minio.Minio, oss *biz.OssUseCase) *OssService {
	return &OssService{task: task, minio: minio, oss: oss}
}
