package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"goflow/internal/model"
	"goflow/internal/pkg/errcode"
	"goflow/internal/pkg/logger"
	"goflow/internal/repository"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProductService interface {
	List(page, size int) ([]model.Product, int64, error)
	GetByID(ctx context.Context, id uint) (*model.Product, error)
	Create(product *model.Product) error
	Update(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, id uint) error
}

type productService struct {
	repo        repository.ProductRepository
	rdb         *redis.Client
	cacheExpire time.Duration
}

func NewProductService(repo repository.ProductRepository, rdb *redis.Client, expire int) ProductService {
	if expire <= 0 {
		expire = 600 // 默认 10 分钟
	}
	return &productService{
		repo:        repo,
		rdb:         rdb,
		cacheExpire: time.Duration(expire) * time.Second,
	}
}

func (s *productService) List(page, size int) ([]model.Product, int64, error) {
	return s.repo.List(page, size)
}

func (s *productService) GetByID(ctx context.Context, id uint) (*model.Product, error) {
	cacheKey := fmt.Sprintf("product:%d", id)
	if s.rdb != nil {
		cached, err := s.rdb.Get(ctx, cacheKey).Result()
		if err == nil {
			var product model.Product
			if json.Unmarshal([]byte(cached), &product) == nil {
				return &product, nil
			}
		} else if err != redis.Nil {
			logger.WithCtx(ctx).Warnw("get product cache failed", "error", err)
		}
	}

	product, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.ErrProductNotFound
		}
		return nil, errcode.ErrInternal
	}

	if s.rdb != nil {
		if data, err := json.Marshal(product); err == nil {
			if err := s.rdb.Set(ctx, cacheKey, data, s.cacheExpire).Err(); err != nil {
				logger.WithCtx(ctx).Warnw("set product cache failed", "error", err)
			}
		}
	}

	return product, nil
}

func (s *productService) Create(product *model.Product) error {
	return s.repo.Create(product)
}

func (s *productService) Update(ctx context.Context, product *model.Product) error {
	if err := s.repo.Update(product); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.ErrProductNotFound
		}
		return errcode.ErrInternal
	}
	if s.rdb != nil {
		cacheKey := fmt.Sprintf("product:%d", product.ID)
		if err := s.rdb.Del(ctx, cacheKey).Err(); err != nil {
			logger.WithCtx(ctx).Warnw("delete product cache failed", "error", err)
		}
	}
	return nil
}

func (s *productService) Delete(ctx context.Context, id uint) error {
	if err := s.repo.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.ErrProductNotFound
		}
		return errcode.ErrInternal
	}
	if s.rdb != nil {
		cacheKey := fmt.Sprintf("product:%d", id)
		if err := s.rdb.Del(ctx, cacheKey).Err(); err != nil {
			logger.WithCtx(ctx).Warnw("delete product cache failed", "error", err)
		}
	}
	return nil
}
