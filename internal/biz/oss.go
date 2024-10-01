package biz

import (
	"context"
	"strconv"
	"strings"

	"github.com/go-cinch/common/constant"
	"github.com/go-cinch/common/copierx"
	"github.com/go-cinch/common/page"
	"github.com/go-cinch/common/utils"
	"github.com/pkg/errors"
	"oss/internal/conf"
)

type CreateOss struct {
	ID   uint64 `json:"id,string"`
	Name string `json:"name"`
}

type Oss struct {
	ID   uint64 `json:"id,string"`
	Name string `json:"name"`
}

type FindOss struct {
	Page page.Page `json:"page"`
	Name *string   `json:"name"`
}

type FindOssCache struct {
	Page page.Page `json:"page"`
	List []Oss     `json:"list"`
}

type UpdateOss struct {
	ID   uint64  `json:"id,string"`
	Name *string `json:"name,omitempty"`
}

type OssRepo interface {
	Create(ctx context.Context, item *CreateOss) error
	Get(ctx context.Context, id uint64) (*Oss, error)
	Find(ctx context.Context, condition *FindOss) []Oss
	Update(ctx context.Context, item *UpdateOss) error
	Delete(ctx context.Context, ids ...uint64) error
}

type OssUseCase struct {
	c     *conf.Bootstrap
	repo  OssRepo
	tx    Transaction
	cache Cache
}

func NewOssUseCase(c *conf.Bootstrap, repo OssRepo, tx Transaction, cache Cache) *OssUseCase {
	return &OssUseCase{
		c:    c,
		repo: repo,
		tx:   tx,
		cache: cache.WithPrefix(strings.Join([]string{
			c.Name, "oss",
		}, "_")),
	}
}

func (uc *OssUseCase) Create(ctx context.Context, item *CreateOss) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) error {
			return uc.repo.Create(ctx, item)
		})
	})
}

func (uc *OssUseCase) Get(ctx context.Context, id uint64) (rp *Oss, err error) {
	rp = &Oss{}
	action := strings.Join([]string{"get", strconv.FormatUint(id, 10)}, "_")
	str, err := uc.cache.Get(ctx, action, func(ctx context.Context) (string, error) {
		return uc.get(ctx, action, id)
	})
	if err != nil {
		return
	}
	utils.Json2Struct(&rp, str)
	if rp.ID == constant.UI0 {
		err = ErrRecordNotFound(ctx)
		return
	}
	return
}

func (uc *OssUseCase) get(ctx context.Context, action string, id uint64) (res string, err error) {
	// read data from db and write to cache
	rp := &Oss{}
	item, err := uc.repo.Get(ctx, id)
	notFound := errors.Is(err, ErrRecordNotFound(ctx))
	if err != nil && !notFound {
		return
	}
	copierx.Copy(&rp, item)
	res = utils.Struct2Json(rp)
	uc.cache.Set(ctx, action, res, notFound)
	return
}

func (uc *OssUseCase) Find(ctx context.Context, condition *FindOss) (rp []Oss, err error) {
	// use md5 string as cache replay json str, key is short
	action := strings.Join([]string{"find", utils.StructMd5(condition)}, "_")
	str, err := uc.cache.Get(ctx, action, func(ctx context.Context) (string, error) {
		return uc.find(ctx, action, condition)
	})
	if err != nil {
		return
	}
	var cache FindOssCache
	utils.Json2Struct(&cache, str)
	condition.Page = cache.Page
	rp = cache.List
	return
}

func (uc *OssUseCase) find(ctx context.Context, action string, condition *FindOss) (res string, err error) {
	// read data from db and write to cache
	list := uc.repo.Find(ctx, condition)
	var cache FindOssCache
	cache.List = list
	cache.Page = condition.Page
	res = utils.Struct2Json(cache)
	uc.cache.Set(ctx, action, res, len(list) == 0)
	return
}

func (uc *OssUseCase) Update(ctx context.Context, item *UpdateOss) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) (err error) {
			err = uc.repo.Update(ctx, item)
			return
		})
	})
}

func (uc *OssUseCase) Delete(ctx context.Context, ids ...uint64) error {
	return uc.tx.Tx(ctx, func(ctx context.Context) error {
		return uc.cache.Flush(ctx, func(ctx context.Context) (err error) {
			err = uc.repo.Delete(ctx, ids...)
			return
		})
	})
}
