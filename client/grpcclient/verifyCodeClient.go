package grpcclient

import (
	"github.com/Diode222/etcd_service_discovery/etcdservice"
	"github.com/DualVectorFoil/Zelda/app/conf"
	pb "github.com/DualVectorFoil/Zelda/protobuf"
	"sync"
)

// TODO 写成一个依赖库，和自己的protobuf放在一起，每次更新都去更新这个依赖库就好了
var verifyCodeClient pb.VerifyCodeServiceClient
var verifyCodeClientOnce sync.Once

func NewVerifyCodeClient() pb.VerifyCodeServiceClient {
	verifyCodeClientOnce.Do(func() {
		verifyCodeClient = etcdservice.NewServiceManager(conf.ETCD_ADDRESS).GetClient(conf.MARIO_SERVICE_ADDR, pb.NewVerifyCodeClientWrapper).(pb.VerifyCodeServiceClient)
	})
	return verifyCodeClient
}
