package biz

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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
	list = make([]OcrItem, len(images))
	for i, image := range images {
		now := time.Now().UnixMilli()
		imageData, err := toBase64(image)
		if err != nil {
			list[i].Msg = err.Error()
			continue
		}
		list[i].Base64Data = imageData
		list[i].ParseLatency = time.Now().UnixMilli() - now
	}

	for i, item := range list {
		if item.Msg != "" {
			continue
		}
		now := time.Now().UnixMilli()
		res, err := uc.ocr.Ocr(ctx, item.Base64Data)
		if err != nil {
			list[i].Msg = err.Error()
			continue
		}
		list[i].OcrLatency = time.Now().UnixMilli() - now
		list[i].Confidence = strconv.FormatFloat(res.Confidence, 'f', -1, 64)
		list[i].Text = res.Text
		boxes, _ := json.Marshal(res.TextRegion)
		list[i].Boxes = string(boxes)
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
