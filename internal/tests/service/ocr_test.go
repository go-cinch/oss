package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-cinch/common/log"
	"github.com/go-cinch/common/utils"
	"github.com/samber/lo"
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

var Images = []string{
	"https://minio-api.go-cinch.top/ocr/1-5803.png",
	"https://minio-api.go-cinch.top/ocr/2-6849.png",
	"https://minio-api.go-cinch.top/ocr/3-9130.png",
	"https://minio-api.go-cinch.top/ocr/4-5774.png",
	"https://minio-api.go-cinch.top/ocr/5-0218.png",
	"https://minio-api.go-cinch.top/ocr/6-2697.png",
	"https://minio-api.go-cinch.top/ocr/7-8800.png",
	"https://minio-api.go-cinch.top/ocr/8-0520.png",
	"https://minio-api.go-cinch.top/ocr/9-7125.png",
	"https://minio-api.go-cinch.top/ocr/10-5505.png",
	"https://minio-api.go-cinch.top/ocr/11-0039.png",
	"https://minio-api.go-cinch.top/ocr/12-5632.png",
	"https://minio-api.go-cinch.top/ocr/13-7511.png",
	"https://minio-api.go-cinch.top/ocr/14-6021.png",
	"https://minio-api.go-cinch.top/ocr/15-0018.png",
	"https://minio-api.go-cinch.top/ocr/16-8551.png",
	"https://minio-api.go-cinch.top/ocr/17-5098.png",
	"https://minio-api.go-cinch.top/ocr/18-1852.png",
	"https://minio-api.go-cinch.top/ocr/19-1021.png",
	"https://minio-api.go-cinch.top/ocr/20-6159.png",
	"https://minio-api.go-cinch.top/ocr/21-5318.png",
	"https://minio-api.go-cinch.top/ocr/22-7268.png",
	"https://minio-api.go-cinch.top/ocr/23-7365.png",
	"https://minio-api.go-cinch.top/ocr/24-7233.png",
	"https://minio-api.go-cinch.top/ocr/25-9654.png",
	"https://minio-api.go-cinch.top/ocr/26-8668.png",
	"https://minio-api.go-cinch.top/ocr/27-5423.png",
	"https://minio-api.go-cinch.top/ocr/28-5180.png",
	"https://minio-api.go-cinch.top/ocr/29-6120.png",
	"https://minio-api.go-cinch.top/ocr/30-5574.png",
	"https://minio-api.go-cinch.top/ocr/31-3243.png",
	"https://minio-api.go-cinch.top/ocr/32-2649.png",
	"https://minio-api.go-cinch.top/ocr/33-9213.png",
	"https://minio-api.go-cinch.top/ocr/34-9226.png",
	"https://minio-api.go-cinch.top/ocr/35-5770.png",
	"https://minio-api.go-cinch.top/ocr/36-5871.png",
	"https://minio-api.go-cinch.top/ocr/37-3771.png",
	"https://minio-api.go-cinch.top/ocr/38-4851.png",
	"https://minio-api.go-cinch.top/ocr/39-1484.png",
	"https://minio-api.go-cinch.top/ocr/40-2170.png",
	"https://minio-api.go-cinch.top/ocr/41-3707.png",
	"https://minio-api.go-cinch.top/ocr/42-7249.png",
	"https://minio-api.go-cinch.top/ocr/43-0207.png",
	"https://minio-api.go-cinch.top/ocr/44-2910.png",
	"https://minio-api.go-cinch.top/ocr/45-0658.png",
	"https://minio-api.go-cinch.top/ocr/46-7826.png",
	"https://minio-api.go-cinch.top/ocr/47-0636.png",
	"https://minio-api.go-cinch.top/ocr/48-1223.png",
	"https://minio-api.go-cinch.top/ocr/49-4512.png",
	"https://minio-api.go-cinch.top/ocr/50-4069.png",
	"https://minio-api.go-cinch.top/ocr/51-7432.png",
	"https://minio-api.go-cinch.top/ocr/52-4711.png",
	"https://minio-api.go-cinch.top/ocr/53-1152.png",
	"https://minio-api.go-cinch.top/ocr/54-4600.png",
	"https://minio-api.go-cinch.top/ocr/55-0003.png",
	"https://minio-api.go-cinch.top/ocr/56-7819.png",
	"https://minio-api.go-cinch.top/ocr/57-4866.png",
	"https://minio-api.go-cinch.top/ocr/58-4238.png",
	"https://minio-api.go-cinch.top/ocr/59-5779.png",
	"https://minio-api.go-cinch.top/ocr/60-4683.png",
	"https://minio-api.go-cinch.top/ocr/61-0038.png",
	"https://minio-api.go-cinch.top/ocr/62-9239.png",
	"https://minio-api.go-cinch.top/ocr/63-1518.png",
	"https://minio-api.go-cinch.top/ocr/64-9863.png",
	"https://minio-api.go-cinch.top/ocr/65-4267.png",
	"https://minio-api.go-cinch.top/ocr/66-7729.png",
	"https://minio-api.go-cinch.top/ocr/67-2734.png",
	"https://minio-api.go-cinch.top/ocr/68-7114.png",
	"https://minio-api.go-cinch.top/ocr/69-4030.png",
	"https://minio-api.go-cinch.top/ocr/70-0542.png",
	"https://minio-api.go-cinch.top/ocr/71-0463.png",
	"https://minio-api.go-cinch.top/ocr/72-1444.png",
	"https://minio-api.go-cinch.top/ocr/73-7206.png",
	"https://minio-api.go-cinch.top/ocr/74-8956.png",
	"https://minio-api.go-cinch.top/ocr/75-0694.png",
	"https://minio-api.go-cinch.top/ocr/76-3355.png",
	"https://minio-api.go-cinch.top/ocr/77-2967.png",
	"https://minio-api.go-cinch.top/ocr/78-6860.png",
	"https://minio-api.go-cinch.top/ocr/79-4682.png",
	"https://minio-api.go-cinch.top/ocr/80-6049.png",
	"https://minio-api.go-cinch.top/ocr/81-2360.png",
	"https://minio-api.go-cinch.top/ocr/82-7469.png",
	"https://minio-api.go-cinch.top/ocr/83-2009.png",
	"https://minio-api.go-cinch.top/ocr/84-6567.png",
	"https://minio-api.go-cinch.top/ocr/85-9017.png",
	"https://minio-api.go-cinch.top/ocr/86-2663.png",
	"https://minio-api.go-cinch.top/ocr/87-9212.png",
	"https://minio-api.go-cinch.top/ocr/88-8142.png",
	"https://minio-api.go-cinch.top/ocr/89-6677.png",
	"https://minio-api.go-cinch.top/ocr/90-3854.png",
	"https://minio-api.go-cinch.top/ocr/91-2531.png",
	"https://minio-api.go-cinch.top/ocr/92-8717.png",
	"https://minio-api.go-cinch.top/ocr/93-2437.png",
	"https://minio-api.go-cinch.top/ocr/94-7978.png",
	"https://minio-api.go-cinch.top/ocr/95-9143.png",
	"https://minio-api.go-cinch.top/ocr/96-4930.png",
	"https://minio-api.go-cinch.top/ocr/97-4646.png",
	"https://minio-api.go-cinch.top/ocr/98-8772.png",
	"https://minio-api.go-cinch.top/ocr/99-5815.png",
	"https://minio-api.go-cinch.top/ocr/100-6249.png",
	"https://minio-api.go-cinch.top/ocr/101-4900.png",
	"https://minio-api.go-cinch.top/ocr/102-4165.png",
	"https://minio-api.go-cinch.top/ocr/103-6347.png",
	"https://minio-api.go-cinch.top/ocr/104-2372.png",
	"https://minio-api.go-cinch.top/ocr/105-5567.png",
	"https://minio-api.go-cinch.top/ocr/106-6190.png",
	"https://minio-api.go-cinch.top/ocr/107-1864.png",
	"https://minio-api.go-cinch.top/ocr/108-4479.png",
	"https://minio-api.go-cinch.top/ocr/109-1348.png",
	"https://minio-api.go-cinch.top/ocr/110-9783.png",
	"https://minio-api.go-cinch.top/ocr/111-2239.png",
	"https://minio-api.go-cinch.top/ocr/112-7707.png",
	"https://minio-api.go-cinch.top/ocr/113-8624.png",
	"https://minio-api.go-cinch.top/ocr/114-6386.png",
	"https://minio-api.go-cinch.top/ocr/115-0537.png",
	"https://minio-api.go-cinch.top/ocr/116-1320.png",
	"https://minio-api.go-cinch.top/ocr/117-1596.png",
	"https://minio-api.go-cinch.top/ocr/118-7273.png",
	"https://minio-api.go-cinch.top/ocr/119-7125.png",
	"https://minio-api.go-cinch.top/ocr/120-9232.png",
	"https://minio-api.go-cinch.top/ocr/121-9067.png",
	"https://minio-api.go-cinch.top/ocr/122-6577.png",
	"https://minio-api.go-cinch.top/ocr/123-7522.png",
	"https://minio-api.go-cinch.top/ocr/124-3093.png",
	"https://minio-api.go-cinch.top/ocr/125-2800.png",
	"https://minio-api.go-cinch.top/ocr/126-8255.png",
	"https://minio-api.go-cinch.top/ocr/127-6055.png",
	"https://minio-api.go-cinch.top/ocr/128-0359.png",
	"https://minio-api.go-cinch.top/ocr/129-8581.png",
	"https://minio-api.go-cinch.top/ocr/130-7744.png",
	"https://minio-api.go-cinch.top/ocr/131-5145.png",
	"https://minio-api.go-cinch.top/ocr/132-1431.png",
	"https://minio-api.go-cinch.top/ocr/133-8316.png",
	"https://minio-api.go-cinch.top/ocr/134-7772.png",
	"https://minio-api.go-cinch.top/ocr/135-6310.png",
	"https://minio-api.go-cinch.top/ocr/136-1112.png",
	"https://minio-api.go-cinch.top/ocr/137-3951.png",
	"https://minio-api.go-cinch.top/ocr/138-8812.png",
	"https://minio-api.go-cinch.top/ocr/139-7167.png",
	"https://minio-api.go-cinch.top/ocr/140-1792.png",
	"https://minio-api.go-cinch.top/ocr/141-2754.png",
	"https://minio-api.go-cinch.top/ocr/142-3304.png",
	"https://minio-api.go-cinch.top/ocr/143-9906.png",
	"https://minio-api.go-cinch.top/ocr/144-8265.png",
	"https://minio-api.go-cinch.top/ocr/145-9465.png",
	"https://minio-api.go-cinch.top/ocr/146-9130.png",
	"https://minio-api.go-cinch.top/ocr/147-2809.png",
	"https://minio-api.go-cinch.top/ocr/148-1547.png",
	"https://minio-api.go-cinch.top/ocr/149-0658.png",
	"https://minio-api.go-cinch.top/ocr/150-0834.png",
	"https://minio-api.go-cinch.top/ocr/151-0752.png",
	"https://minio-api.go-cinch.top/ocr/152-1744.png",
	"https://minio-api.go-cinch.top/ocr/153-7361.png",
	"https://minio-api.go-cinch.top/ocr/154-2390.png",
	"https://minio-api.go-cinch.top/ocr/155-5495.png",
	"https://minio-api.go-cinch.top/ocr/156-1532.png",
	"https://minio-api.go-cinch.top/ocr/157-2971.png",
	"https://minio-api.go-cinch.top/ocr/158-5496.png",
	"https://minio-api.go-cinch.top/ocr/159-7755.png",
	"https://minio-api.go-cinch.top/ocr/160-6465.png",
	"https://minio-api.go-cinch.top/ocr/161-0953.png",
	"https://minio-api.go-cinch.top/ocr/162-8608.png",
	"https://minio-api.go-cinch.top/ocr/163-2651.png",
	"https://minio-api.go-cinch.top/ocr/164-5028.png",
	"https://minio-api.go-cinch.top/ocr/165-1907.png",
	"https://minio-api.go-cinch.top/ocr/166-6761.png",
	"https://minio-api.go-cinch.top/ocr/167-8267.png",
	"https://minio-api.go-cinch.top/ocr/168-3363.png",
	"https://minio-api.go-cinch.top/ocr/169-1167.png",
	"https://minio-api.go-cinch.top/ocr/170-1093.png",
	"https://minio-api.go-cinch.top/ocr/171-8993.png",
	"https://minio-api.go-cinch.top/ocr/172-1809.png",
	"https://minio-api.go-cinch.top/ocr/173-2071.png",
	"https://minio-api.go-cinch.top/ocr/174-4139.png",
	"https://minio-api.go-cinch.top/ocr/175-4477.png",
	"https://minio-api.go-cinch.top/ocr/176-8348.png",
	"https://minio-api.go-cinch.top/ocr/177-6802.png",
	"https://minio-api.go-cinch.top/ocr/178-0199.png",
	"https://minio-api.go-cinch.top/ocr/179-1239.png",
	"https://minio-api.go-cinch.top/ocr/180-5514.png",
	"https://minio-api.go-cinch.top/ocr/181-5489.png",
	"https://minio-api.go-cinch.top/ocr/182-3414.png",
	"https://minio-api.go-cinch.top/ocr/183-2093.png",
	"https://minio-api.go-cinch.top/ocr/184-9336.png",
	"https://minio-api.go-cinch.top/ocr/185-1132.png",
	"https://minio-api.go-cinch.top/ocr/186-0751.png",
	"https://minio-api.go-cinch.top/ocr/187-8923.png",
	"https://minio-api.go-cinch.top/ocr/188-0583.png",
	"https://minio-api.go-cinch.top/ocr/189-6358.png",
	"https://minio-api.go-cinch.top/ocr/190-9532.png",
	"https://minio-api.go-cinch.top/ocr/191-6977.png",
	"https://minio-api.go-cinch.top/ocr/192-8182.png",
	"https://minio-api.go-cinch.top/ocr/193-9188.png",
	"https://minio-api.go-cinch.top/ocr/194-5293.png",
	"https://minio-api.go-cinch.top/ocr/195-5223.png",
	"https://minio-api.go-cinch.top/ocr/196-5857.png",
	"https://minio-api.go-cinch.top/ocr/197-8010.png",
	"https://minio-api.go-cinch.top/ocr/198-4346.png",
	"https://minio-api.go-cinch.top/ocr/199-3511.png",
	"https://minio-api.go-cinch.top/ocr/200-7665.png",
}

func TestOssService_BatchOcr(t *testing.T) {
	// m, err := minio.New(
	// 	minio.WithEndpoint("minio-api.go-cinch.top"),
	// 	minio.WithKey("5dPKP49qZrm5i4875zUM"),
	// 	minio.WithSecret("Y0lyCQ8gp2HpZRXUeIr04LRYih6uWQEQF5A0LGAe"),
	// 	minio.WithBucket("ocr"),
	// 	minio.WithSSL(true),
	// )
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	//
	// memory := base64Captcha.NewMemoryStore(100, 10*time.Minute)
	// num := 4
	// width := num * 45
	//
	// files := make([]string, 0)
	// for i := 0; i < 200; i++ {
	// 	c := base64Captcha.NewCaptcha(base64Captcha.NewDriverDigit(80, width, num, 0.7, 80), memory)
	// 	_, img, answer, _ := c.Generate()
	// 	img = strings.TrimPrefix(img, "data:image/png;base64,")
	// 	imageData, _ := base64.StdEncoding.DecodeString(img)
	// 	object := strings.Join([]string{"val", strconv.FormatInt(int64(i+1), 10) + "-" + answer + ".png"}, "/")
	// 	tmpPath := strings.Join([]string{"/tmp", object}, "/")
	// 	_ = os.MkdirAll(path.Dir(tmpPath), os.ModePerm)
	// 	_ = os.WriteFile(tmpPath, imageData, 0644)
	// 	files = append(files, tmpPath)
	// }
	//
	// images := make([]string, 0)
	// for _, file := range files {
	// 	object := path.Base(file)
	// 	_, err = m.FPutObject(context.Background(), object, file)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	images = append(images, fmt.Sprintf("https://minio-api.go-cinch.top/ocr/%s", object))
	// }
	// fmt.Println(utils.Struct2Json(images))

	for {
		go func() {
			_ = callOcr(context.Background(), lo.Shuffle(Images)[:5])
		}()
		time.Sleep(1000*time.Millisecond)
	}
}

func callOcr(ctx context.Context, images []string) (err error) {
	uri := "https://app.go-cinch.top/api/oss/ocr"
	log.
		WithContext(ctx).
		WithField("uri", uri).
		Info("call ocr api")

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewReader([]byte(utils.Struct2Json(
		struct {
			List    []string `json:"list"`
			Latency string   `json:"latency"`
		}{
			List: images,
		},
	))))

	req.Header.Add("X-Canary",
		"iwantsit")
	req.Header.Add("Accept",
		"application/json")
	req.Header.Add("Content-Type",
		"application/json")
	req.Header.Add("Authorization",
		"Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiODlIRUsyOFkiLCJleHAiOjE3MzA5NDYzOTMsInBsYXRmb3JtIjoicGMifQ.DaVxfcHe9gkmpn5fD1bSrUrkMK0XPOGxkZ-f_0wXPjGcTCOHEZeKv5y74LK7K6S1z2oUzXReweSaoccTOMLiag")

	body, err := call(ctx, req)
	if err != nil {
		return
	}

	var response struct {
		List    []map[string]interface{} `json:"list"`
		Latency string                   `json:"latency"`
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
	fmt.Println(response.Latency)
	return
}

func call(ctx context.Context, req *http.Request) (data []byte, err error) {
	client := &http.Client{
		Timeout: time.Duration(600) * time.Second,
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
