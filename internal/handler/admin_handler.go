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

type adminLoginService interface {
	Login(ctx context.Context, username, password string) (*response.LoginResult, error)
}

type AdminHandler struct {
	adminSvc adminLoginService
}

func NewAdminHandler(adminSvc adminLoginService) *AdminHandler {
	return &AdminHandler{adminSvc: adminSvc}
}

// Login 管理员登录
func (h *AdminHandler) Login(c *gin.Context) {
	var r req.AdminLoginReq
	if err := c.ShouldBindJSON(&r); err != nil {
		response.ErrorWithMsg(c, 400, errcode.CodeBadRequest, validator.Translate(err, i18n.GetLang(c)))
		return
	}

	result, err := h.adminSvc.Login(c.Request.Context(), r.Username, r.Password)
	if err != nil {
		writeServiceError(c, err)
		return
	}

	response.Success(c, result)
}
