package dao

import "context"

type IVerifyCodeDao interface {
	SetVerifyCodeInfo(ctx context.Context, phoneNum string, verifyCode string) error
	IsVerifyCodeAvailable(ctx context.Context, phoneNum string, verifyCode string) (bool, error)
}
