package admin

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	emailRegex       = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	phoneNumberRegex = regexp.MustCompile(`^\+?[0-9]{10,15}$`)
)

func (r CreateAdminAccount) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required.Error("Email is required"),
			validation.By(func(value interface{}) error {
				email, _ := value.(string)
				if !emailRegex.MatchString(email) {
					return validation.NewError("validation_email", "Invalid email format")
				}
				return nil
			}),
		),
		validation.Field(&r.Password, validation.Required.Error("Password is required")),
		validation.Field(&r.Role, validation.Required.Error("Role is required")),
	)
}

func (r LoginAdminAccount) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required.Error("Email is required"),
			validation.By(func(value interface{}) error {
				email, _ := value.(string)
				if !emailRegex.MatchString(email) {
					return validation.NewError("validation_email", "Invalid email format")
				}
				return nil
			}),
		),
		validation.Field(&r.Password, validation.Required.Error("Password is required")),
	)
}
