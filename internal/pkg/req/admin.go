package req

// AdminLoginReq 管理员登录请求
type AdminLoginReq struct {
	Username string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}
