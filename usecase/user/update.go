package user

type UpdateRequest struct {
	Name                 string `json:"name" validate:"max=50"`
	Email                string `json:"email" validate:"omitempty,max=255,email"`
	Password             string `json:"password" validate:"omitempty,min=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"omitempty,min=6"`
}
