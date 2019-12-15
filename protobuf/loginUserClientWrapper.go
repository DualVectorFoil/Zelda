package protobuf

import "google.golang.org/grpc"

func NewLoginUserClientWrapper(cc *grpc.ClientConn) interface{} {
	return NewLoginUserServiceClient(cc)
}
