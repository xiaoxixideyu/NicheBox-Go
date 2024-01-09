// Code generated by goctl. DO NOT EDIT.
package types

type CheckEmailExistsReqeust struct {
	Email string `json:"email"`
}

type CheckEmailExistsResponse struct {
	Exist bool `json:"exist"`
}

type CheckVerificationCodeCriticalUserInfoRequest struct {
	Email string `json:"email"`
	Code  string `json:"code""`
}

type CheckVerificationCodeCriticalUserInfoResponse struct {
}

type ForgetPasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password""`
	Code        string `json:"code""`
}

type ForgetPasswordResponse struct {
}

type LoginReqeust struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenReqeust struct {
}

type RefreshTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type RegisterResponse struct {
	LoginSuccess bool   `json:"login_success"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type SendVerificationCodeCriticalUserInfoRequest struct {
	Destination string `json:"destination"`
}

type SendVerificationCodeCriticalUserInfoResponse struct {
}

type SendVerificationCodePWDRequest struct {
	Destination string `json:"destination"`
}

type SendVerificationCodePWDResponse struct {
}

type SendVerificationCodeRegisterRequest struct {
	Destination string `json:"destination"`
}

type SendVerificationCodeRegisterResponse struct {
}

type SetCriticalUserInfoRequest struct {
	Password string `json:"password""`
}

type SetCriticalUserInfoResponse struct {
}
