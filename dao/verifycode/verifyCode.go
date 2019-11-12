package verifycode

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/DualVectorFoil/Zelda/protobuf"
	"github.com/DualVectorFoil/Zelda/utils/ptr"
)

type VerifyCode struct {
	Client pb.VerifyCodeServiceClient
}

func NewVerifyCode(client pb.VerifyCodeServiceClient) *VerifyCode {
	return &VerifyCode{
		Client: client,
	}
}

func (v *VerifyCode) SetVerifyCodeInfo(ctx context.Context, phoneNum string, verifyCode string) error {
	verifyCodeInfo := &pb.VerifyCodeInfo{
		PhoneNum:             ptr.StringPtr(phoneNum),
		VerifyCode:           ptr.StringPtr(verifyCode),
	}
	resp, err := v.Client.SetVerifyCodeInfo(ctx, verifyCodeInfo)
	if err != nil {
		fmt.Println("SetVerifyCodeInfo failed.")
		return err
	}

	if !resp.GetStatus() {
		fmt.Println("SetVerifyCodeInfo failed.")
		return errors.New("SetVerifyCodeInfo failed.")
	}

	return nil
}

func (v *VerifyCode) IsVerifyCodeAvailable(ctx context.Context, phoneNum string, verifyCode string) (bool, error) {
	verifyCodeInfo := &pb.VerifyCodeInfo{
		PhoneNum:             ptr.StringPtr(phoneNum),
		VerifyCode:           ptr.StringPtr(verifyCode),
	}
	resp, err := v.Client.IsVerifyCodeAvailable(ctx, verifyCodeInfo)
	if err != nil {
		fmt.Println("get IsVerifyCodeAvailable failed.")
		return false, err
	}

	return resp.GetStatus(), nil
}
