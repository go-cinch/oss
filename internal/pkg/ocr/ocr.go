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

type Req struct {
	Image              string  `json:"image"`
	BoxThreshold       *string `json:"boxThreshold"`       // only for helmet, default 0.5
	DetMethod          *string `json:"detMethod"`          // only for helmet, default auto
	MinSliceWidth      *string `json:"minSliceWidth"`      // only for helmet, default 800
	MinSliceHeight     *string `json:"minSliceHeight"`     // only for helmet, default 800
	SliceWidth         *string `json:"sliceWidth"`         // only for helmet, default 640
	SliceHeight        *string `json:"sliceHeight"`        // only for helmet, default 640
	OverlapThreshold   *string `json:"overlapThreshold"`   // only for helmet, default 0.6
	OverlapWidthRatio  *string `json:"overlapWidthRatio"`  // only for helmet, default 0.2
	OverlapHeightRatio *string `json:"overlapHeightRatio"` // only for helmet, default 0.2
}

type Resp struct {
	List []Detail `json:"list"`
}

type Detail struct {
	Original     string  `json:"original"`
	Msg          string  `json:"msg"`
	ParseLatency int     `json:"parseLatency"`
	DetLatency   int     `json:"detLatency"`
	ClsLatency   int     `json:"clsLatency"`
	RecLatency   int     `json:"recLatency"`
	Points       []Point `json:"points"`
}

type Point struct {
	Confidence string `json:"confidence"`
	Text       string `json:"text"`
	TextRegion string `json:"textRegion"`
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

func (c Ocr) OcrPredict(ctx context.Context, condition *Req) (rp *Resp, err error) {
	host := strings.Join([]string{c.host, "predict/ocr"}, "/")
	res, err := callAPI[*Resp](
		ctx,
		request{
			host: host,
			body: bytes.NewReader([]byte(utils.Struct2Json(struct {
				Images []string `json:"images"`
			}{
				Images: []string{condition.Image},
			}))),
		},
	)
	if err != nil {
		return
	}
	if len(res.List) > 0 {
		rp = res
		return
	}
	err = errors.New("cannot recognize")
	return
}

func (c Ocr) HelmetPredict(ctx context.Context, condition *Req) (rp *Resp, err error) {
	host := strings.Join([]string{c.host, "predict/helmet"}, "/")
	if condition.BoxThreshold == nil {
		condition.BoxThreshold = &[]string{"0.5"}[0]
	}
	if condition.DetMethod == nil {
		condition.DetMethod = &[]string{"auto"}[0]
	}
	if condition.MinSliceWidth == nil {
		condition.MinSliceWidth = &[]string{"800"}[0]
	}
	if condition.MinSliceHeight == nil {
		condition.MinSliceHeight = &[]string{"800"}[0]
	}
	if condition.SliceWidth == nil {
		condition.SliceWidth = &[]string{"640"}[0]
	}
	if condition.SliceHeight == nil {
		condition.SliceHeight = &[]string{"640"}[0]
	}
	if condition.OverlapThreshold == nil {
		condition.OverlapThreshold = &[]string{"0.6"}[0]
	}
	if condition.OverlapWidthRatio == nil {
		condition.OverlapWidthRatio = &[]string{"0.2"}[0]
	}
	if condition.OverlapHeightRatio == nil {
		condition.OverlapHeightRatio = &[]string{"0.2"}[0]
	}
	res, err := callAPI[*Resp](
		ctx,
		request{
			host: host,
			body: bytes.NewReader([]byte(utils.Struct2Json(struct {
				Images             []string `json:"images"`
				BoxThreshold       string   `json:"box_threshold"`
				DetMethod          string   `json:"det_method"`
				MinSliceWidth      string   `json:"min_slice_width"`
				MinSliceHeight     string   `json:"min_slice_height"`
				SliceWidth         string   `json:"slice_width"`
				SliceHeight        string   `json:"slice_height"`
				OverlapThreshold   string   `json:"overlap_threshold"`
				OverlapWidthRatio  string   `json:"overlap_width_ratio"`
				OverlapHeightRatio string   `json:"overlap_height_ratio"`
			}{
				Images:             []string{condition.Image},
				BoxThreshold:       *condition.BoxThreshold,
				DetMethod:          *condition.DetMethod,
				MinSliceWidth:      *condition.MinSliceWidth,
				MinSliceHeight:     *condition.MinSliceHeight,
				SliceWidth:         *condition.SliceWidth,
				SliceHeight:        *condition.SliceHeight,
				OverlapThreshold:   *condition.OverlapThreshold,
				OverlapWidthRatio:  *condition.OverlapWidthRatio,
				OverlapHeightRatio: *condition.OverlapHeightRatio,
			}))),
		},
	)
	if err != nil {
		return
	}
	if len(res.List) > 0 {
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

func callAPI[T any](ctx context.Context, params request) (data T, err error) {
	uri := params.host
	log.
		WithContext(ctx).
		Info("call ocr api: %s", uri)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, uri, params.body)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	body, err := call(ctx, req)
	if err != nil {
		return
	}

	var response T
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
	data = response
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
