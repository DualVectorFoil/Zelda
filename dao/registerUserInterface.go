package dao

import (
	"context"
)

type IRegisterUserDao interface {
	RegisterUser(ctx context.Context, phoneNum string, userName string, pwd string, verifyCode string) error
}
