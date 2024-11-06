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

func (uc *OcrUseCase) Ocr(ctx context.Context, images ...string) (list []OcrItem) {
	var wg sync.WaitGroup
	concurrency := uc.c.Ocr.Concurrency
	if concurrency <= 0 {
		concurrency = 10
	}
	sem := make(chan struct{}, concurrency)
	list = make([]OcrItem, len(images))
	for i, item := range images {
		wg.Add(1)
		go func(idx int, data string) {
			defer wg.Done()
			list[idx] = *uc.processOcrRequest(ctx, sem, data)
			list[idx].Original = data
		}(i, item)
	}
	wg.Wait()
	return
}

func (uc *OcrUseCase) processOcrRequest(ctx context.Context, sem chan struct{}, image string) (item *OcrItem) {
	sem <- struct{}{}
	defer func() { <-sem }()
	item = &OcrItem{}
	now := time.Now().UnixMilli()
	imageData, err := toBase64(image)
	item.ParseLatency = time.Now().UnixMilli() - now
	if err != nil {
		item.Msg = err.Error()
		return
	}
	now = time.Now().UnixMilli()
	res, err := uc.ocr.Ocr(ctx, imageData)
	item.OcrLatency = time.Now().UnixMilli() - now
	if err != nil {
		item.Msg = err.Error()
		return
	}
	item.Confidence = strconv.FormatFloat(res.Confidence, 'f', -1, 64)
	item.Text = res.Text
	boxes, _ := json.Marshal(res.TextRegion)
	item.Boxes = string(boxes)
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
