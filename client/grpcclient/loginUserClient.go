package grpcclient

import (
	"github.com/Diode222/etcd_service_discovery/etcdservice"
	"github.com/DualVectorFoil/Zelda/app/conf"
	pb "github.com/DualVectorFoil/Zelda/protobuf"
	"sync"
)

var loginUserClient pb.LoginUserServiceClient
var loginUserClientOnce sync.Once

func NewLoginUserClient() pb.LoginUserServiceClient {
	loginUserClientOnce.Do(func() {
		loginUserClient = etcdservice.NewServiceManager(conf.ETCD_ADDRESS).GetClient(conf.MARIO_SERVICE_NAME, pb.NewLoginUserClientWrapper).(pb.LoginUserServiceClient)
	})
	return loginUserClient
}
