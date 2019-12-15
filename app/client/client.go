package client

import "github.com/DualVectorFoil/Zelda/client/grpcclient"

var LoginUserClient = grpcclient.NewLoginUserClient()

var VerifyCodeClient = grpcclient.NewVerifyCodeClient()

var RegisterUserClient = grpcclient.NewRegisterUserClient()
