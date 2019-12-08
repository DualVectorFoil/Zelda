package protobuf

import "google.golang.org/grpc"

func NewVerifyCodeClientWrapper(cc *grpc.ClientConn) interface{} {
	return NewVerifyCodeServiceClient(cc)
}
