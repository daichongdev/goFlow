package handler

import (
	"context"

	"gonio/internal/pkg/errcode"
	"gonio/internal/pkg/i18n"
	"gonio/internal/pkg/req"
	"gonio/internal/pkg/response"
	"gonio/internal/pkg/validator"

	"github.com/gin-gonic/gin"
)

type userLoginService interface {
	Login(ctx context.Context, username, password string) (*response.LoginResult, error)
}

type UserHandler struct {
	userSvc userLoginService
}

func NewUserHandler(userSvc userLoginService) *UserHandler {
	return &UserHandler{userSvc: userSvc}
}

// Login 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	var r req.LoginReq
	if err := c.ShouldBindJSON(&r); err != nil {
		response.ErrorWithMsg(c, 400, errcode.CodeBadRequest, validator.Translate(err, i18n.GetLang(c)))
		return
	}

	result, err := h.userSvc.Login(c.Request.Context(), r.Username, r.Password)
	if err != nil {
		writeServiceError(c, err)
		return
	}

	response.Success(c, result)
}
