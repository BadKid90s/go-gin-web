package request

type RegisterRequest struct {
	Username  *string `json:"username"`
	LoginName *string `json:"loginName"`
	Password  *string `json:"password"`
}

type LoginRequest struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}
