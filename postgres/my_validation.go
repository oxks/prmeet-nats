package postgres

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (s UserSignupParams) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Nickname, validation.Required, validation.Length(1, 20)),
		validation.Field(&s.Password, validation.Required, validation.Length(1, 250)),
		validation.Field(&s.Email, validation.Required, is.Email),
	)
}
