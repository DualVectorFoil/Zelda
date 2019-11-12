package handler

import (
	"Zelda/controller"
	"sync"
)

type handler struct {
	UserCtrl *controller.UserCtrl
}

var h *handler
var handlerOnce sync.Once

func GetHandlerInstance() *handler {
	handlerOnce.Do(func() {
		h = &handler{
			UserCtrl: controller.NewUserCtrl(),
		}
	})
	return h
}
