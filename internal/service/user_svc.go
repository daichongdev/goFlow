package service

import (
	"errors"
	"goflow/internal/config"
	"goflow/internal/pkg/response"
	"time"

	"goflow/internal/model"
	"goflow/internal/pkg/errcode"
	"goflow/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(username, password, nickname string) error
	Login(username, password string) (*response.LoginResult, error)
}

type userService struct {
	repo      repository.UserRepository
	jwtSecret []byte
	jwtExpire time.Duration
}

func NewUserService(repo repository.UserRepository, cfg *config.Config) UserService {
	jwtExpire := cfg.JWT.Expire
	if jwtExpire <= 0 {
		jwtExpire = 7200
	}
	return &userService{
		repo:      repo,
		jwtSecret: []byte(cfg.JWT.Secret),
		jwtExpire: time.Duration(jwtExpire) * time.Second,
	}
}

// CreateUser 创建用户
func (s *userService) CreateUser(username, password, nickname string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errcode.ErrInternal
	}
	return s.repo.Create(&model.User{
		Username: username,
		Password: string(hashed),
		Nickname: nickname,
		Status:   1,
	})
}

func (s *userService) Login(username, password string) (*response.LoginResult, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.ErrUserOrPassword
		}
		return nil, errcode.ErrInternal
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errcode.ErrUserOrPassword
	}

	// 检查用户状态
	if user.Status != 1 {
		return nil, errcode.ErrUserDisabled
	}

	// 生成 JWT
	expireAt := time.Now().Add(s.jwtExpire)
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     "user",
		"exp":      expireAt.Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return nil, errcode.ErrInternal
	}

	return &response.LoginResult{
		Token:    tokenStr,
		ExpireAt: expireAt.Unix(),
		User: gin.H{
			"id":       user.ID,
			"username": user.Username,
			"nickname": user.Nickname,
			"avatar":   user.Avatar,
		},
	}, nil
}
