package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-cinch/common/log"
	"github.com/go-cinch/common/minio"
	"github.com/go-cinch/common/utils"
	"github.com/mojocn/base64Captcha"
	"oss/api/oss"
	"oss/internal/tests/mock"
)

func TestOssService_Ocr(t *testing.T) {
	s := mock.OssService()
	res, err := s.Ocr(context.Background(), &oss.OcrRequest{
		List: []string{
			"https://minio-api.go-cinch.top/ocr/1.png",
			"https://minio-api.go-cinch.top/ocr/2.png",
			"https://minio-api.go-cinch.top/ocr/3.png",
			"https://minio-api.go-cinch.top/ocr/4.png",
			"https://minio-api.go-cinch.top/ocr/5.png",
			"https://minio-api.go-cinch.top/ocr/6.png",
			"https://minio-api.go-cinch.top/ocr/7.png",
			"https://minio-api.go-cinch.top/ocr/8.png",
			"https://minio-api.go-cinch.top/ocr/9.png",
			"https://minio-api.go-cinch.top/ocr/10.png",
			"https://minio-api.go-cinch.top/ocr/11.png",
			"https://minio-api.go-cinch.top/ocr/12.png",
			"https://minio-api.go-cinch.top/ocr/13.png",
			"https://minio-api.go-cinch.top/ocr/14.png",
			"https://minio-api.go-cinch.top/ocr/15.png",
			"https://minio-api.go-cinch.top/ocr/16.png",
			"https://minio-api.go-cinch.top/ocr/17.png",
			"https://minio-api.go-cinch.top/ocr/18.png",
			"https://minio-api.go-cinch.top/ocr/19.png",
			"https://minio-api.go-cinch.top/ocr/20.png",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

func TestOssService_BatchOcr(t *testing.T) {
	m, err := minio.New(
		minio.WithEndpoint("minio-api.go-cinch.top"),
		minio.WithKey("5dPKP49qZrm5i4875zUM"),
		minio.WithSecret("Y0lyCQ8gp2HpZRXUeIr04LRYih6uWQEQF5A0LGAe"),
		minio.WithBucket("ocr"),
		minio.WithSSL(true),
	)
	if err != nil {
		t.Error(err)
		return
	}

	memory := base64Captcha.NewMemoryStore(100, 10*time.Minute)
	num := 4
	width := num * 45
	c := base64Captcha.NewCaptcha(base64Captcha.NewDriverDigit(80, width, num, 0.7, 80), memory)
	_, img, answer, _ := c.Generate()
	img = strings.TrimPrefix(img, "data:image/png;base64,")
	imageData, _ := base64.StdEncoding.DecodeString(img)

	files := make([]string, 0)
	for i := 0; i < 100; i++ {
		object := strings.Join([]string{"val", strconv.FormatInt(int64(i+1), 10) + "-" + answer + ".png"}, "/")
		tmpPath := strings.Join([]string{"/tmp", object}, "/")
		_ = os.MkdirAll(path.Dir(tmpPath), os.ModePerm)
		_ = os.WriteFile(tmpPath, imageData, 0644)
		files = append(files, tmpPath)
	}

	images := make([]string, 0)
	for _, file := range files {
		object := path.Base(file)
		_, err = m.FPutObject(context.Background(), object, file)
		if err != nil {
			continue
		}
		images = append(images, fmt.Sprintf("https://minio-api.go-cinch.top/ocr/%s", object))
	}
	_ = callOcr(context.Background(), images)
}

func callOcr(ctx context.Context, images []string) (err error) {
	uri := "https://app.go-cinch.top/api/oss/ocr"
	log.
		WithContext(ctx).
		WithField("uri", uri).
		Info("call ocr api")

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewReader([]byte(utils.Struct2Json(
		struct {
			List []string `json:"list"`
		}{
			List: images,
		},
	))))

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiODlIRUsyOFkiLCJleHAiOjE3MzA5NDYzOTMsInBsYXRmb3JtIjoicGMifQ.DaVxfcHe9gkmpn5fD1bSrUrkMK0XPOGxkZ-f_0wXPjGcTCOHEZeKv5y74LK7K6S1z2oUzXReweSaoccTOMLiag")

	body, err := call(ctx, req)
	if err != nil {
		return
	}

	var response struct {
		List []map[string]interface{} `json:"list"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		msg := "invalid body"
		log.
			WithContext(ctx).
			WithError(err).
			Warn(msg)
		err = errors.New(msg)
		return
	}
	fmt.Println(response.List)
	return
}

func call(ctx context.Context, req *http.Request) (data []byte, err error) {
	client := &http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		if isTimeout(err) {
			log.
				WithContext(ctx).
				Warn("connect timeout")
			return
		}
		log.
			WithContext(ctx).
			WithError(err).
			Warn("connect failed")
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.
			WithContext(ctx).
			WithError(err).
			Warn("read body failed")
		return
	}

	if res.StatusCode != http.StatusOK {
		msg := strings.Join([]string{"invalid status code: ", strconv.Itoa(res.StatusCode)}, "")
		log.
			WithContext(ctx).
			WithField("http.error", res.Status).
			WithField("http.body", string(body)).
			Warn(msg)
		err = errors.New(msg)
		return
	}
	data = body
	return
}

func isTimeout(err error) bool {
	var neterr net.Error
	if errors.As(err, &neterr) {
		return neterr != nil && neterr.Timeout()
	}
	return false
}
