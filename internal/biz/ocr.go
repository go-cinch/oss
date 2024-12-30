package biz

import (
	"context"
	"sync"

	"github.com/go-cinch/common/copierx"
	"github.com/go-cinch/common/log"
	"oss/api/oss"
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

func (uc *OcrUseCase) Ocr(ctx context.Context, req *oss.OcrRequest) (list []ocr.Detail) {
	var wg sync.WaitGroup
	concurrency := uc.c.Ocr.Concurrency
	if concurrency <= 0 {
		concurrency = 10
	}
	sem := make(chan struct{}, concurrency)
	list = make([]ocr.Detail, len(req.List))
	for i, item := range req.List {
		wg.Add(1)
		go func(idx int, data string) {
			defer wg.Done()
			list[idx] = *uc.processOcrRequest(ctx, sem, req.Category, &ocr.Req{
				Image:              data,
				BoxThreshold:       req.BoxThreshold,
				DetMethod:          req.DetMethod,
				MinSliceWidth:      req.MinSliceWidth,
				MinSliceHeight:     req.MinSliceHeight,
				SliceWidth:         req.SliceWidth,
				SliceHeight:        req.SliceHeight,
				OverlapThreshold:   req.OverlapThreshold,
				OverlapWidthRatio:  req.OverlapWidthRatio,
				OverlapHeightRatio: req.OverlapHeightRatio,
			})
		}(i, item)
	}
	wg.Wait()
	return
}

func (uc *OcrUseCase) processOcrRequest(ctx context.Context, sem chan struct{}, category *oss.OcrCategory, condition *ocr.Req) (rp *ocr.Detail) {
	sem <- struct{}{}
	defer func() { <-sem }()
	rp = &ocr.Detail{
		Original: condition.Image,
	}
	var err error
	var res *ocr.Resp
	if category == nil {
		category = &[]oss.OcrCategory{oss.OcrCategory_OCR}[0]
	}
	switch *category {
	case oss.OcrCategory_OCR:
		res, err = uc.ocr.OcrPredict(ctx, condition)
	case oss.OcrCategory_HELMET:
		res, err = uc.ocr.HelmetPredict(ctx, condition)
	default:
		log.WithContext(ctx).Warn("invalid category: %d", category)
		return
	}
	if err != nil {
		rp.Msg = err.Error()
		return
	}
	copierx.Copy(rp, res.List[0])
	return
}
