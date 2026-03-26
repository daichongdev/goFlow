package svc

import (
	"goflow/internal/config"
	"goflow/internal/mq"
	"goflow/internal/repository"
	"goflow/internal/service"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ServiceContext 集中管理所有依赖，作为唯一的依赖注入容器。
// 新增模块只需在此处添加字段并在 NewServiceContext 中初始化。
type ServiceContext struct {
	Config *config.Config

	// 基础设施
	DB    *gorm.DB
	Redis *redis.Client

	// Repository 层
	ProductRepo repository.ProductRepository
	UserRepo    repository.UserRepository
	AdminRepo   repository.AdminRepository

	// Service 层
	ProductSvc service.ProductService
	UserSvc    service.UserService
	AdminSvc   service.AdminService

	// MQ
	MQPublisher *mq.Publisher
}

// NewServiceContext 根据配置和基础设施连接创建 ServiceContext，
// 完成 repo → service 的依赖接线。
func NewServiceContext(cfg *config.Config, db *gorm.DB, rdb *redis.Client, mqPublisher *mq.Publisher) *ServiceContext {
	// Repository
	productRepo := repository.NewProductRepo(db)
	userRepo := repository.NewUserRepo(db)
	adminRepo := repository.NewAdminRepo(db)

	// Service
	productSvc := service.NewProductService(productRepo, rdb, cfg.Server.CacheExpire)
	userSvc := service.NewUserService(userRepo, cfg)
	adminSvc := service.NewAdminService(adminRepo, cfg)

	return &ServiceContext{
		Config: cfg,

		DB:    db,
		Redis: rdb,

		ProductRepo: productRepo,
		UserRepo:    userRepo,
		AdminRepo:   adminRepo,

		ProductSvc: productSvc,
		UserSvc:    userSvc,
		AdminSvc:   adminSvc,

		MQPublisher: mqPublisher,
	}
}
