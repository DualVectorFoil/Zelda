package model

type ProfileModel struct {
	PhoneNum     string `json:"phone_num"`
	AvatarUrl    string `json:"avatar_url"`
	UserName     string `json:"user_name"`
	Locale       string `json:"locale"`
	Bio          string `json:"bio"`
	Followers    int32  `json:"followers"`
	Following    int32  `json:"following"`
	ArtworkCount int32  `json:"artwork_count"`
	PWD          string `json:"pwd"`
	RegisteredAt int64  `json:"registered_at"`
	LastLoginAt  int64  `json:"last_login_at"`
	Token        string `json:"token"`
}
