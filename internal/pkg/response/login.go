package response

// LoginResult 登录返回结果
type LoginResult struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expire_at"`
	User     any    `json:"user"`
}
