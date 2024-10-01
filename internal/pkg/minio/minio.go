package minio

import (
	"errors"

	"github.com/go-cinch/common/log"
	"github.com/go-cinch/common/minio"
	"github.com/google/wire"
	"oss/internal/conf"
)

// ProviderSet is minio providers.
var ProviderSet = wire.NewSet(New)

// New is initialize minio from config
func New(c *conf.Bootstrap) (m *minio.Minio, err error) {
	m, err = minio.New(
		minio.WithEndpoint(c.Minio.Endpoint),
		minio.WithKey(c.Minio.Key),
		minio.WithSecret(c.Minio.Secret),
		minio.WithBucket(c.Minio.Bucket),
		minio.WithExpire(c.Minio.Expire),
		minio.WithSSL(c.Minio.Ssl),
	)
	if err != nil {
		log.Error(err)
		err = errors.New("initialize minio failed")
		return
	}

	log.Info("initialize minio success")
	return
}
