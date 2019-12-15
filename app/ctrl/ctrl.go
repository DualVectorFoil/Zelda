package ctrl

import (
	"github.com/DualVectorFoil/Zelda/app/dao"
	"github.com/DualVectorFoil/Zelda/controller"
)

var UserCtrl = controller.NewUserCtrl(dao.LoginUserDao, dao.RegisterUserDao, dao.VerifyCodeDao)
