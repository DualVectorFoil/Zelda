package verifycode

import (
	"context"
	"errors"
	pb "github.com/DualVectorFoil/Zelda/protobuf"
	"github.com/DualVectorFoil/Zelda/utils/ptr"
	"github.com/sirupsen/logrus"
	"time"
)

const MAX_REQUEST_TIME  = time.Second * 3

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

	timeoutCtx, _ := context.WithTimeout(ctx, MAX_REQUEST_TIME)

	resp, err := v.Client.SetVerifyCodeInfo(timeoutCtx, verifyCodeInfo)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"phoneNum": phoneNum,
			"verifyCode": verifyCode,
			"err": err.Error(),
		}).Error("SetVerifyCodeInfo failed.")
		return err
	}

	if !resp.GetStatus() {
		logrus.WithFields(logrus.Fields{
			"phoneNum": phoneNum,
			"verifyCode": verifyCode,
		}).Error("SetVerifyCodeInfo failed.")
		return errors.New("SetVerifyCodeInfo failed.")
	}

	return nil
}

func (v *VerifyCode) IsVerifyCodeAvailable(ctx context.Context, phoneNum string, verifyCode string) (bool, error) {
	verifyCodeInfo := &pb.VerifyCodeInfo{
		PhoneNum:             ptr.StringPtr(phoneNum),
		VerifyCode:           ptr.StringPtr(verifyCode),
	}

	timeoutCtx, _ := context.WithTimeout(ctx, MAX_REQUEST_TIME)

	resp, err := v.Client.IsVerifyCodeAvailable(timeoutCtx, verifyCodeInfo)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"phoneNum": phoneNum,
			"verifyCode": verifyCode,
			"err": err.Error(),
		}).Error("get IsVerifyCodeAvailable failed.")
		return false, err
	}

	return resp.GetStatus(), nil
}
