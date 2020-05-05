package request

// 用户注册
type RegisterStruct struct {
	Email    string `json:"email"`
	NickName string `json:"nickName"`
	Password string `json:"password"`
}

type LoginStruct struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
