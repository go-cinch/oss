package biz

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

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

func (uc *OcrUseCase) Ocr(ctx context.Context, images ...string) (list []OcrResult) {
	var wg sync.WaitGroup
	concurrency := uc.c.Ocr.Concurrency
	if concurrency <= 0 {
		concurrency = 10
	}
	sem := make(chan struct{}, concurrency)
	list = make([]OcrResult, len(images))
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

func (uc *OcrUseCase) processOcrRequest(ctx context.Context, sem chan struct{}, image string) (rp *OcrResult) {
	sem <- struct{}{}
	defer func() { <-sem }()
	rp = &OcrResult{
		Original: image,
	}
	now := time.Now().UnixMilli()
	imageData, err := toBase64(image)
	rp.ParseLatency = time.Now().UnixMilli() - now
	if err != nil {
		rp.Msg = err.Error()
		return
	}
	now = time.Now().UnixMilli()
	res, err := uc.ocr.Ocr(ctx, imageData)
	rp.OcrLatency = time.Now().UnixMilli() - now
	if err != nil {
		rp.Msg = err.Error()
		return
	}
	for _, v1 := range res {
		for _, v2 := range v1 {
			boxes, _ := json.Marshal(v2.TextRegion)
			rp.List = append(rp.List, OcrItem{
				Boxes:      string(boxes),
				Confidence: strconv.FormatFloat(v2.Confidence, 'f', -1, 64),
				Text:       v2.Text,
			})
		}
	}
	return
}

func toBase64(image string) (string, error) {
	resp, err := http.Get(image)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	base64Image := base64.StdEncoding.EncodeToString(imageData)
	return base64Image, nil
}
