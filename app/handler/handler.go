package handler

import (
	"github.com/DualVectorFoil/Zelda/app/ctrl"
	"github.com/DualVectorFoil/Zelda/controller"
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
			UserCtrl: ctrl.UserCtrl,
		}
	})
	return h
}
