package grpcclient

import (
	"github.com/Diode222/etcd_service_discovery/etcdservice"
	"github.com/DualVectorFoil/Zelda/app/conf"
	pb "github.com/DualVectorFoil/Zelda/protobuf"
	"sync"
)

var registerUserClient pb.RegisterUserServiceClient
var registerUserClientOnce sync.Once

func NewRegisterUserClient() pb.RegisterUserServiceClient {
	registerUserClientOnce.Do(func() {
		registerUserClient = etcdservice.NewServiceManager(conf.ETCD_ADDRESS).GetClient(conf.MARIO_SERVICE_NAME, pb.NewRegisterUserClientWrapper).(pb.RegisterUserServiceClient)
	})
	return registerUserClient
}
