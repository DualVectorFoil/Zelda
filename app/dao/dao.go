package dao

import (
	"github.com/DualVectorFoil/Zelda/app/client"
	"github.com/DualVectorFoil/Zelda/dao/registerUser"
	"github.com/DualVectorFoil/Zelda/dao/verifycode"
)

var VerifyCodeDao = verifycode.NewVerifyCodeDao(client.VerifyCodeClient)

var RegisterUserDao = registerUser.NewRegisterUserDao(client.RegisterUserClient)
