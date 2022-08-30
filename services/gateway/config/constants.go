package config

var Bucket = []string{"background", "routine", "avatar", "avatar-account"}

const (
	ForgotPassword          = "Instructions sent! Please check your email."
	ChangePassword          = "Password has been changed."
	UpdateAccount           = "Account has been updated."
	DeactiveAccount         = "Account has been deactivated."
	ActiveAccount           = "Account has been activated."
	DeleteSuccess           = "Delete success."
	Page                    = 1
	PerPage                 = 20
	RoutineBucketName       = "routine"
	AvatarAccountBucketName = "avatar-account"
	BackgroundBucketName    = "background"
	AvatarBucketName        = "avatar"
)

var (
	Auth_message_code = map[string]int32{
		"ok":          0,
		"unspecified": 1,
		"rpc error: code = Unknown desc = wrong email or password": 2,
		"rpc error: code = Unknown desc = email does not exist":    3,
	}
)
