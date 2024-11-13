package ocr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-cinch/common/log"
	"github.com/go-cinch/common/utils"
	"github.com/google/wire"
	"oss/internal/conf"
)

var ProviderSet = wire.NewSet(New)

type Resp struct {
	Confidence float64 `json:"confidence"`
	Text       string  `json:"text"`
	TextRegion [][]int `json:"text_region"`
}

type Ocr struct {
	host string
}

func New(c *conf.Bootstrap) (api API, err error) {
	log.Info("initialize ocr success")
	return &Ocr{
		host: c.Ocr.Host,
	}, nil
}

func (c Ocr) Ocr(ctx context.Context, image string) (rp [][]Resp, err error) {
	host := strings.Join([]string{c.host, "ocr_system"}, "/")
	res, err := callOcr[[][]Resp](
		ctx,
		request{
			host: host,
			body: bytes.NewReader([]byte(utils.Struct2Json(struct {
				Images []string `json:"images"`
			}{
				Images: []string{image},
			}))),
		},
	)
	if err != nil {
		return
	}
	if len(res) > 0 {
		rp = res
		return
	}
	err = errors.New("cannot recognize")
	return
}

type request struct {
	host string
	body io.Reader
}

type resp[T any] struct {
	Status  string `json:"status"`
	Msg     string `json:"msg"`
	Results T      `json:"results"`
}

func callOcr[T any](ctx context.Context, params request) (data T, err error) {
	uri := params.host
	log.
		WithContext(ctx).
		WithField("uri", uri).
		Info("call ocr api")

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, uri, params.body)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	body, err := call(ctx, req)
	if err != nil {
		return
	}

	var response resp[T]
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
	if response.Status != "000" {
		msg := response.Msg
		log.
			WithContext(ctx).
			Warn(msg)
		err = errors.New(msg)
		return
	}
	data = response.Results
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
