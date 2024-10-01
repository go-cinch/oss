package service

import (
	"context"
	"strconv"
	"strings"
	"testing"

	"github.com/go-cinch/common/proto/params"
	"github.com/google/uuid"
	"oss/api/oss"
	"oss/internal/tests/mock"
)

func TestOssService_CreateOss(t *testing.T) {
	s := mock.OssService()
	ctx := context.Background()
	userID := uuid.NewString()
	ctx = mock.NewContextWithUserId(ctx, userID)

	_, err := s.CreateOss(ctx, &oss.CreateOssRequest{
		Name: "oss1",
	})
	if err != nil {
		t.Error(err)
		return
	}
	_, err = s.CreateOss(ctx, &oss.CreateOssRequest{
		Name: "oss2",
	})
	if err != nil {
		t.Error(err)
		return
	}
	res1, _ := s.FindOss(ctx, &oss.FindOssRequest{
		Page: &params.Page{
			Disable: true,
		},
	})
	if res1 == nil || len(res1.List) != 2 {
		t.Error("res1 len must be 2")
		return
	}
	res2, err := s.GetOss(ctx, &oss.GetOssRequest{
		Id: res1.List[0].Id,
	})
	if err != nil {
		t.Error(err)
		return
	}
	if res2.Name != res1.List[0].Name {
		t.Errorf("res2.Name must be %s", res1.List[0].Name)
		return
	}
	_, err = s.DeleteOss(ctx, &params.IdsRequest{
		Ids: strings.Join([]string{
			strconv.FormatUint(res1.List[0].Id, 10),
			strconv.FormatUint(res1.List[1].Id, 10),
		}, ","),
	})
	if err != nil {
		t.Error(err)
		return
	}
	res3, _ := s.FindOss(ctx, &oss.FindOssRequest{
		Page: &params.Page{
			Disable: true,
		},
	})
	if res3 == nil || len(res3.List) != 0 {
		t.Error("res3 len must be 0")
		return
	}
}
