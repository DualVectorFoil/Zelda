package dao

import (
	"github.com/DualVectorFoil/Zelda/app/client"
	"github.com/DualVectorFoil/Zelda/dao/login"
	"github.com/DualVectorFoil/Zelda/dao/registerUser"
	"github.com/DualVectorFoil/Zelda/dao/verifycode"
)

var LoginUserDao = login.NewLoginUserDao(client.LoginUserClient)

var VerifyCodeDao = verifycode.NewVerifyCodeDao(client.VerifyCodeClient)

var RegisterUserDao = registerUser.NewRegisterUserDao(client.RegisterUserClient)
