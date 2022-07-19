package auth

type SignUpRequest struct {
	Email                string `json:"email" validate:"required,max=255,email"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=6"`
}
