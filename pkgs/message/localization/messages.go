package localization

// sucess message

const (
	// customer account register sucess messages
	MsgSuccessCreateCustomer = "Customer created successfully"
	MsgSuccessVerifyCustomer = "Account verified successfully"

	// customer account login messages
	MsgSuccessLoginCustomer = "Login successful"
)

// error messages
const (
	// basic error messages
	MsgInvalidInput   = "Invalid input"
	MsgInvalidKey     = "Invalid key"
	MsgInvalidEncData = "Invalid encryption data"
	MsgInvalidPadding = "Invalid padding"

	// System error messages
	MsgInternalServerError  = "Internal server error occurred"
	MsgServiceUnavailable   = "Service is temporarily unavailable"
	MsgDatabaseError        = "Database operation failed"
	MsgNetworkError         = "Network error occurred"
	MsgTimeoutError         = "Request timeout occurred"
	MsgUnexpectedError      = "An unexpected error occurred"
	MsgConfigurationError   = "Configuration error"
	MsgExternalServiceError = "External service error"
	MsgValidationFailed     = "Validation failed"
	MsgUserForbidden        = "User access is forbidden"
	MsgUserUnauthorized     = "User is unauthorized"

	//
	MsgValidationPassed  = "Validation passed successfully"
	MsgHealthCheckPassed = "Health check passed"
	MsgHealthCheckFailed = "Health check failed"
	MsgBadRequest        = "Bad request"
	// customer account register error messages
	MsgCreateCustomer                    = "Failed to create customer"
	MsgVerifiedFailed                    = "Failed to verify customer"
	MsgCreatedCustomerEmailAlreadyExists = "Customer email already exists"
	MsgEmailDuplicate                    = "Customer email already exists"
	MsgOtpExpired                        = "OTP has expired"
	MsgInvalidOtp                        = "Invalid OTP provided provided by customer"
	MsgPassword                          = "Password is too short"
	MsgPasswordNotMatch                  = "Passwords do not match"
	MsgLoginCustomer                     = "Failed to login customer"
	MsgUserNotFound                      = "User not found"
	MsgCustomerNotActive                 = "Customer is not active"
	MsgInvalidCredentials                = "Invalid email or password"
)
