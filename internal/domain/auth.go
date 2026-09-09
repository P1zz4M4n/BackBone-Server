package domain

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token         string `json:"token"`
	RefereshToken string `json:"refresh_token"`
	ExpireAt      int64  `json:"expire_at"`
}

type RegisterInput struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordInput struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

type AuthUseCase interface {
	RegisterUser(input *RegisterInput) (*TokenResponse, error)
	Login(input *LoginInput) (*TokenResponse, error)
	Logout(token string) error
	RefreshToken(refreshToken string) (*TokenResponse, error)
	ResetPassword(input *ResetPasswordInput) error
	SendOTP(email string) error
	VerifyOTP(email, otp string) error
}
