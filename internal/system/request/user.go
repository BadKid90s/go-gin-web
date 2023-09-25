package request

type RegisterRequest struct {
	UserName  *string `json:"userName"`
	LoginName *string `json:"loginName"`
	Password  *string `json:"password"`
}

type LoginRequest struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}
