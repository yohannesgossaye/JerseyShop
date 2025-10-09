package customeraccount

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	emailRegex       = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	phoneNumberRegex = regexp.MustCompile(`^\+?[0-9]{10,15}$`)
)

func (r CreateCustomer) Validator() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.FirstName, validation.Required.Error("First name is required")),
		validation.Field(&r.LastName, validation.Required.Error("Last name is required")),
		validation.Field(&r.Email, validation.Required.Error("Email is required"),
			validation.By(func(value interface{}) error {
				email, _ := value.(string)
				if !emailRegex.MatchString(email) {
					return validation.NewError("validation_email", "Invalid email format")
				}
				return nil
			}),
		),
		validation.Field(&r.PhoneNumber, validation.Required.Error("Phone number is required"),
			validation.By(func(value interface{}) error {
				phoneNumber, _ := value.(string)
				if !phoneNumberRegex.MatchString(phoneNumber) {
					return validation.NewError("validation_phone_number", "Invalid phone number format")
				}
				return nil
			})),
		validation.Field(&r.Password, validation.Required.Error("Password is required")),
		validation.Field(&r.Country, validation.Required.Error("Country is required")),
		validation.Field(&r.City, validation.Required.Error("City is required")),
		validation.Field(&r.Address, validation.Required.Error("Address is required")),
		validation.Field(&r.ZipCode, validation.Required.Error("Zip code is required")),
		validation.Field(&r.Gender, validation.Required.Error("Gender is required")),
		validation.Field(&r.DateOfBirth, validation.Required.Error("Date of birth is required")),
	)
}

func (r LoginCustomer) Validator() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required.Error("Email is required")),
		validation.Field(&r.Password, validation.Required.Error("Password is required")),
	)
}

func (r VerifyCustomerOtp) Validator() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required.Error("Email is required")),
		validation.Field(&r.OTPCode, validation.Required.Error("OTP code is required")),
	)
}
