package dto

import "github.com/asaskevich/govalidator"

type UserLoginDto struct {
	Username string `valid:"required,length(8|20)"`
	Password string `valid:"required,length(8|20)"`
}

func demo() {
	govalidator.IsURL(`http://user@pass:domain.com/path/page`)
}
