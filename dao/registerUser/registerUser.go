package registerUser

import (
	"context"
	"errors"
	pb "github.com/DualVectorFoil/Zelda/protobuf"
	"github.com/DualVectorFoil/Zelda/utils/ptr"
	"time"
)

const MAX_REQUEST_TIME = time.Second * 10

type RegisterUserDao struct {
	Client pb.RegisterUserServiceClient
}

func NewRegisterUserDao(client pb.RegisterUserServiceClient) *RegisterUserDao {
	return &RegisterUserDao{Client: client}
}

func (r *RegisterUserDao) RegisterUser(ctx context.Context, phoneNum string, userName string, pwd string, verifyCode string) error {
	registerInfo := &pb.RegisterInfo{
		PhoneNum:   ptr.StringPtr(phoneNum),
		UserName:   ptr.StringPtr(userName),
		Pwd:        ptr.StringPtr(pwd),
		VerifyCode: ptr.StringPtr(verifyCode),
	}

	timeoutCtx, _ := context.WithTimeout(ctx, MAX_REQUEST_TIME)

	resp, err := r.Client.RegisterUser(timeoutCtx, registerInfo)
	if resp != nil && !resp.GetStatus() {
		err = errors.New(resp.GetErr())
	}
	return err
}
