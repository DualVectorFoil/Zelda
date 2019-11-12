package dao

import (
	"github.com/DualVectorFoil/Zelda/app/client"
	"github.com/DualVectorFoil/Zelda/dao/verifycode"
)

var VerifyCodeDao = verifycode.NewVerifyCode(client.MarioClient)
