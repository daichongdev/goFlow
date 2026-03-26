package errcode

import "fmt"

// AppError 统一业务错误类型
type AppError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	httpStatus int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func (e *AppError) HTTPStatus() int {
	return e.httpStatus
}

// New 创建自定义错误
func New(code int, httpStatus int, msg string) *AppError {
	return &AppError{Code: code, httpStatus: httpStatus, Message: msg}
}

// 通用错误码 10001-19999
var (
	ErrBadRequest   = &AppError{Code: 10001, Message: "请求参数错误", httpStatus: 400}
	ErrUnauthorized = &AppError{Code: 10002, Message: "未授权", httpStatus: 401}
	ErrForbidden    = &AppError{Code: 10003, Message: "禁止访问", httpStatus: 403}
	ErrNotFound     = &AppError{Code: 10004, Message: "资源不存在", httpStatus: 404}
	ErrInternal     = &AppError{Code: 10005, Message: "服务器内部错误", httpStatus: 500}
)

// 用户错误码 20001-20099
var (
	ErrUserOrPassword  = &AppError{Code: 20001, Message: "用户名或密码错误", httpStatus: 401}
	ErrUserDisabled    = &AppError{Code: 20002, Message: "用户已被禁用", httpStatus: 403}
	ErrAdminOrPassword = &AppError{Code: 20003, Message: "管理员账号或密码错误", httpStatus: 401}
	ErrAdminDisabled   = &AppError{Code: 20004, Message: "管理员账号已被禁用", httpStatus: 403}
)

// 商品错误码 20101-20199
var (
	ErrProductNotFound = &AppError{Code: 20101, Message: "商品不存在", httpStatus: 404}
	ErrProductOffShelf = &AppError{Code: 20102, Message: "商品已下架", httpStatus: 400}
	ErrStockNotEnough  = &AppError{Code: 20103, Message: "库存不足", httpStatus: 400}
)
