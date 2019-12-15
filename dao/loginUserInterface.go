package dao

import (
	"context"
	"github.com/DualVectorFoil/Zelda/model"
)

type ILoginUserDao interface {
	LoginUserWithPWD(ctx context.Context, userNameInfo string, pwd string) (*model.ProfileModel, error)
	LoginUserWithToken(ctx context.Context, token string) (*model.ProfileModel, error)
}
