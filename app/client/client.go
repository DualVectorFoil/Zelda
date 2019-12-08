package client

import "github.com/DualVectorFoil/Zelda/client/grpcclient"

var VerifyCodeClient = grpcclient.NewVerifyCodeClient()

var RegisterUserClient = grpcclient.NewRegisterUserClient()
