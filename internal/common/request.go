package common

type PageReq struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum"`
}

type PageRest[T any] struct {
	Total int `json:"total"`
	List  []T `json:"list"`
}
