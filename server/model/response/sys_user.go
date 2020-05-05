package response

/**
登录用户
*/
type LoginResult struct {
	Token    string `json:"token"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
