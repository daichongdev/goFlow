package migration

import (
	"gorm.io/gorm"
	"goflow/internal/model"
	"goflow/internal/pkg/logger"
)

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Admin{},
		&model.Product{},
	)
	if err != nil {
		logger.Log.Fatalf("auto migrate failed: %v", err)
	}
	logger.Log.Info("database migration completed")
}
