package response

type UserInfoResponse struct {
	UserName  *string `json:"userName"`
	LoginName *string `json:"loginName"`
	Mobile    *string `json:"mobile"`
	Email     *string `json:"email"`
}
