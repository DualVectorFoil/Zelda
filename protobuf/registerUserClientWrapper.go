package protobuf

import "google.golang.org/grpc"

func NewRegisterUserClientWrapper(cc *grpc.ClientConn) interface{} {
	return NewRegisterUserServiceClient(cc)
}
