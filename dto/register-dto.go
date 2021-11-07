package dto


type RegisterDTO struct {
	Name string `json:"name" form:"name" binding:"required" validate:"min:2"`
	Email string `json:"email" form:"email" binding:"required,email"  validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
}