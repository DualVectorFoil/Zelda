package login

import (
	"context"
	"errors"
	"github.com/DualVectorFoil/Zelda/model"
	pb "github.com/DualVectorFoil/Zelda/protobuf"
	"github.com/DualVectorFoil/Zelda/utils/ptr"
	"time"
)

const MAX_REQUEST_TIME = time.Second * 6

type LoginUserDao struct {
	Client pb.LoginUserServiceClient
}

func NewLoginUserDao(c pb.LoginUserServiceClient) *LoginUserDao {
	return &LoginUserDao{Client: c}
}

func (l *LoginUserDao) LoginUserWithPWD(ctx context.Context, userNameInfo string, pwd string) (*model.ProfileModel, error) {
	loginInfo := &pb.LoginInfo{
		UserNameInfo: ptr.StringPtr(userNameInfo),
		Pwd:          ptr.StringPtr(pwd),
		Token:        ptr.StringPtr(""),
	}

	timeoutCtx, _ := context.WithTimeout(ctx, MAX_REQUEST_TIME)

	resp, err := l.Client.LoginUserWithPWD(timeoutCtx, loginInfo)
	if err != nil {
		return nil, err
	}

	profileInfo := resp.GetProfileInfo()
	if !resp.GetStatus() || profileInfo == nil {
		return nil, errors.New("Login failed.")
	}

	return transformProfileInfo(profileInfo), nil
}

func (l *LoginUserDao) LoginUserWithToken(ctx context.Context, token string) (*model.ProfileModel, error) {
	tokenInfo := &pb.LoginInfo{
		UserNameInfo: ptr.StringPtr(""),
		Pwd:          ptr.StringPtr(""),
		Token:        ptr.StringPtr(token),
	}

	timeoutCtx, _ := context.WithTimeout(ctx, MAX_REQUEST_TIME)

	resp, err := l.Client.LoginUserWithToken(timeoutCtx, tokenInfo)
	if err != nil {
		return nil, err
	}

	profileInfo := resp.GetProfileInfo()
	if !resp.GetStatus() || profileInfo == nil {
		return nil, errors.New("Uncorrected login info, login failed.")
	}

	return transformProfileInfo(profileInfo), nil
}

func transformProfileInfo(profileInfoPB *pb.ProfileInfo) *model.ProfileModel {
	return &model.ProfileModel{
		PhoneNum:     profileInfoPB.GetPhoneNum(),
		AvatarUrl:    profileInfoPB.GetAvatarUrl(),
		UserName:     profileInfoPB.GetUserName(),
		Locale:       profileInfoPB.GetLocale(),
		Bio:          profileInfoPB.GetBio(),
		Followers:    profileInfoPB.GetFollowers(),
		Following:    profileInfoPB.GetFollowing(),
		ArtworkCount: profileInfoPB.GetArtworkCount(),
		PWD:          profileInfoPB.GetPwd(),
		RegisteredAt: profileInfoPB.GetRegisteredAt(),
		LastLoginAt:  profileInfoPB.GetLastLoginAt(),
		Token:        profileInfoPB.GetToken(),
	}
}
