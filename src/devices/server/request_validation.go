package server


import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"

	"github.com/PalmaPedro/tania-core/src/devices/domain"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// REQUIRED ...
const (
	required       = "REQUIRED"
	alphanumeric   = "ALPHANUMERIC"
	alpha          = "ALPHA"
	numeric        = "NUMERIC"
	float          = "FLOAT"
	parseFailed    = "PARSE_FAILED"
	invalidOption  = "INVALID_OPTION"
	notFound       = "NOT_FOUND"
)

// RequestValidation sanitizes request inputs and convert the input to its correct data type.
// This is mostly used to prevent issues like invalid data type or potential SQL Injection.
// So we can focus on processing data without converting data type after this sanitizing.
// This validation doesn't aim to validate business process.
// The business process validation will be handled in each entity's behaviour.
type RequestValidation struct {
}

// RequestValidationError contains fields used for JSON error response
type RequestValidationError struct {
	FieldName    string `json:"field_name"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func (rve RequestValidationError) Error() string {
	return fmt.Sprintf(
		"Field Name: %s, Error Code: %s, Error Message: %s",
		rve.FieldName,
		rve.ErrorCode,
		rve.ErrorMessage,
	)
}

// Message translates error code to meaningful message
func Message(errorCode string) string {
	switch errorCode {
	case required:
		return "This field is required"
	case alphanumeric:
		return "Alphanumeric only"
	case alpha:
		return "Alphabet only"
	case numeric:
		return "Number only"
	case float:
		return "Float only"
	case parseFailed:
		return "Parsing failed. Make sure the input is correct."
	case invalidOption:
		return "This value is not available in options. Please give the correct options."
	case notFound:
		return "Data not found."
	default:
		return "Internal server error"
	}
}

// NewRequestValidationError initializes new RequestValidation struct
func NewRequestValidationError(errorCode, fieldName string) RequestValidationError {
	return RequestValidationError{
		FieldName:    fieldName,
		ErrorCode:    errorCode,
		ErrorMessage: Message(errorCode),
	}
}

// Error wraps errors from application layer and domain layer
// to some format in JSON for response
func Error(c echo.Context, err error) error {
	errorResponse := map[string]string{
		"field_name":    "",
		"error_code":    "",
		"error_message": "",
	}

	file, line := getFileAndLineNumber()

	logData := log.WithFields(log.Fields{
		"user_uid":      c.Get("USER_UID"),
		"request_id":    c.Response().Header().Get(echo.HeaderXRequestID),
		"file":          file,
		"line":          line,
		"error_message": "",
		"field_name":    "",
	})

	if re, ok := err.(domain.DeviceError); ok {
		errorResponse["error_code"] = strconv.Itoa(re.Code)
		errorResponse["error_message"] = re.Error()

		logData.WithField("error_message", re.Error()).Info()

		return c.JSON(http.StatusBadRequest, errorResponse)
	} else if rve, ok := err.(RequestValidationError); ok {
		errorResponse["field_name"] = rve.FieldName
		errorResponse["error_code"] = rve.ErrorCode
		errorResponse["error_message"] = rve.ErrorMessage

		logData.WithField("error_message", rve.ErrorMessage)
		logData.WithField("field_name", rve.FieldName)
		logData.Info()

		return c.JSON(http.StatusBadRequest, rve)
	}

	errorResponse["error_message"] = err.Error()
	logData.WithField("error_message", err.Error()).Error()

	return c.JSON(http.StatusInternalServerError, errorResponse)
}

func getFileAndLineNumber() (string, int) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}

	return file, line
}