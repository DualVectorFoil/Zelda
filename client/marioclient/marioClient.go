package marioclient

import (
	"github.com/Diode222/etcd_service_discovery/etcdservice"
	"github.com/DualVectorFoil/Zelda/app/conf"
	pb "github.com/DualVectorFoil/Zelda/protobuf"
	"sync"
)

// TODO 写成一个依赖库，和自己的protobuf放在一起
var marioClient pb.VerifyCodeServiceClient
var marioClientOnce sync.Once

func NewMarioClient() pb.VerifyCodeServiceClient {
	marioClientOnce.Do(func() {
		marioClient = etcdservice.NewServiceManager(conf.ETCD_ADDRESS).GetClient(conf.MARIO_SERVICE_ADDR, pb.NewVerifyCodeClientWrapper).(pb.VerifyCodeServiceClient)
	})
	return marioClient
}
