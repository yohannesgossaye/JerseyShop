package localization

import "time"

type ResponseCode struct {
	Code       string    `json:"code"`
	StatusCode int       `json:"status_code"`
	Timestamp  time.Time `json:"timestamp"`
	Message    string    `json:"message"`
	Type       string    `json:"type"`
}

func (r ResponseCode) Error() string {
	return r.Message
}

var ResponseCodesList = []ResponseCode{
	// sucess codes
	SucessCustomerCreated,
	SucessCustomerVerified,
	SucessCustomerLogin,

	// error codes
	ErrCustomerNotFound,
	ErrCustomerNotVerified,
	ErrCustomerAlreadyVerified,
	ErrCustomerLogin,
	ErrCustomerPassword,
	ErrCustomerCreate,
	ErrInvalidOtp,
	ErrNotCorrectOtp,
	ErrOtpExpired,
	ErrCustomerCreateDecodeFailed,
	ErrCustomerNotActive,
	ErrInvalidCredentials,
	ErrCustomerEmailAlreadyExists,
}

var (
	// sucess codes
	SucessCustomerCreated = ResponseCode{
		Code:       "SUCCESS_CUSTOMER_CREATED",
		StatusCode: StatusCreated,
		Message:    MsgSuccessCreateCustomer,
		Type:       "success",
	}

	SucessCustomerVerified = ResponseCode{
		Code:       "SUCCESS_CUSTOMER_VERIFIED",
		StatusCode: StatusOK,
		Message:    MsgSuccessVerifyCustomer,
		Type:       "success",
	}
	SucessCustomerLogin = ResponseCode{
		Code:       "SUCCESS_CUSTOMER_LOGIN",
		StatusCode: StatusOK,
		Message:    MsgSuccessLoginCustomer,
		Type:       "success",
	}

	// error codes
	ErrCustomerNotFound = ResponseCode{
		Code:       "ERROR_CUSTOMER_NOT_FOUND",
		StatusCode: StatusNotFound,
		Message:    MsgUserNotFound,
		Type:       "error",
	}
	ErrCustomerNotVerified = ResponseCode{
		Code:       "ERROR_CUSTOMER_NOT_VERIFIED",
		StatusCode: StatusNotFound,
		Message:    MsgVerifiedFailed,
		Type:       "error",
	}
	ErrCustomerAlreadyVerified = ResponseCode{
		Code:       "ERROR_CUSTOMER_ALREADY_VERIFIED",
		StatusCode: StatusConflict,
		Message:    MsgCustomerAlreadyVerified,
		Type:       "error",
	}
	ErrCustomerLogin = ResponseCode{
		Code:       "ERROR_CUSTOMER_LOGIN",
		StatusCode: StatusNotFound,
		Message:    MsgLoginCustomer,
		Type:       "error",
	}
	ErrCustomerPassword = ResponseCode{
		Code:       "ERROR_CUSTOMER_PASSWORD",
		StatusCode: StatusNotFound,
		Message:    MsgPasswordNotMatch,
		Type:       "error",
	}
	ErrCustomerCreate = ResponseCode{
		Code:       "ERROR_CUSTOMER_CREATE",
		StatusCode: StatusNotFound,
		Message:    MsgCreateCustomer,
		Type:       "error",
	}
	ErrCustomerEmailDuplicate = ResponseCode{
		Code:       "ERROR_CUSTOMER_EMAIL_DUPLICATE",
		StatusCode: StatusNotFound,
		Message:    MsgEmailDuplicate,
		Type:       "error",
	}
	ErrInvalidOtp = ResponseCode{
		Code:       "ERROR_INVALID_OTP",
		StatusCode: StatusNotFound,
		Message:    MsgInvalidOtp,
		Type:       "error",
	}
	ErrNotCorrectOtp = ResponseCode{
		Code:       "ERROR_NOT_CORRECT_OTP",
		StatusCode: StatusNotFound,
		Message:    MsgNotCorrectOtp,
		Type:       "error",
	}
	ErrOtpExpired = ResponseCode{
		Code:       "ERROR_OTP_EXPIRED",
		StatusCode: StatusNotFound,
		Message:    MsgOtpExpired,
		Type:       "error",
	}
	ErrCustomerCreateDecodeFailed = ResponseCode{
		Code:       "ERROR_CUSTOMER_CREATE_DECODE_FAILED",
		StatusCode: StatusNotFound,
		Message:    MsgCreateCustomer,
		Type:       "error",
	}
	ErrCustomerNotActive = ResponseCode{
		Code:       "ERROR_CUSTOMER_NOT_ACTIVE",
		StatusCode: StatusNotFound,
		Message:    MsgCustomerNotActive,
		Type:       "error",
	}
	ErrInvalidCredentials = ResponseCode{
		Code:       "ERROR_INVALID_CREDENTIALS",
		StatusCode: StatusNotFound,
		Message:    MsgInvalidCredentials,
		Type:       "error",
	}
	ErrCustomerEmailAlreadyExists = ResponseCode{
		Code:       "ERROR_CUSTOMER_EMAIL_ALREADY_EXISTS",
		StatusCode: StatusConflict,
		Message:    MsgCustomerEmailAlreadyExists,
		Type:       "error",
	}

	//
	ErrorDatabaseError = ResponseCode{
		Code:       "ERROR_DATABASE_ERROR",
		StatusCode: StatusInternalServerError,
		Message:    MsgDatabaseError,
		Type:       "error",
	}

	ErrorNetworkError = ResponseCode{
		Code:       "ERROR_NETWORK_ERROR",
		StatusCode: StatusInternalServerError,
		Message:    MsgNetworkError,
		Type:       "error",
	}

	ErrorTimeoutError = ResponseCode{
		Code:       "ERROR_TIMEOUT_ERROR",
		StatusCode: StatusRequestTimeout,
		Message:    MsgTimeoutError,
		Type:       "error",
	}

	ErrorUnexpectedError = ResponseCode{
		Code:       "ERROR_UNEXPECTED_ERROR",
		StatusCode: StatusInternalServerError,
		Message:    MsgUnexpectedError,
		Type:       "error",
	}

	ErrorConfigurationError = ResponseCode{
		Code:       "ERROR_CONFIGURATION_ERROR",
		StatusCode: StatusInternalServerError,
		Message:    MsgConfigurationError,
		Type:       "error",
	}

	ErrorExternalServiceError = ResponseCode{
		Code:       "ERROR_EXTERNAL_SERVICE_ERROR",
		StatusCode: StatusInternalServerError,
		Message:    MsgExternalServiceError,
		Type:       "error",
	}
	ErrorValidationFailed = ResponseCode{
		Code:       "ERROR_VALIDATION_FAILED",
		StatusCode: StatusBadRequest,
		Message:    MsgValidationFailed,
		Type:       "error",
	}
)
