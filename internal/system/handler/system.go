package handler

type SystemHandler struct {
	User UserHandler
}

func NewSystemHandler(user UserHandler) SystemHandler {
	return SystemHandler{User: user}
}
