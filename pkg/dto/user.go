package dto

type UserLoginDto struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
