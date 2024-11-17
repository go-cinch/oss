package biz

import (
	"context"
	"sync"

	"github.com/go-cinch/common/copierx"
	"oss/internal/conf"
	"oss/internal/pkg/ocr"
)

type OcrUseCase struct {
	c   *conf.Bootstrap
	ocr ocr.API
}

func NewOcrUseCase(c *conf.Bootstrap, ocr ocr.API) *OcrUseCase {
	return &OcrUseCase{
		c:   c,
		ocr: ocr,
	}
}

func (uc *OcrUseCase) Ocr(ctx context.Context, images ...string) (list []ocr.Detail) {
	var wg sync.WaitGroup
	concurrency := uc.c.Ocr.Concurrency
	if concurrency <= 0 {
		concurrency = 10
	}
	sem := make(chan struct{}, concurrency)
	list = make([]ocr.Detail, len(images))
	for i, item := range images {
		wg.Add(1)
		go func(idx int, data string) {
			defer wg.Done()
			list[idx] = *uc.processOcrRequest(ctx, sem, data)
		}(i, item)
	}
	wg.Wait()
	return
}

func (uc *OcrUseCase) processOcrRequest(ctx context.Context, sem chan struct{}, image string) (rp *ocr.Detail) {
	sem <- struct{}{}
	defer func() { <-sem }()
	rp = &ocr.Detail{
		Original: image,
	}
	res, err := uc.ocr.PredictAll(ctx, image)
	if err != nil {
		rp.Msg = err.Error()
		return
	}
	copierx.Copy(rp, res.List[0])
	return
}
