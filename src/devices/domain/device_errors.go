package domain

const (
	// DeviceErrorTitleEmptyCode is used
	DeviceErrorTitleEmptyCode = iota

	// DeviceErrorIDInvalidCode error
	DeviceErrorIDInvalidCode

	// DeviceErrorDescriptionEmptyCode Errors
	DeviceErrorDescriptionEmptyCode

	/*
	// Date Errors
	TaskErrorDueDateEmptyCode
	TaskErrorDueDateInvalidCode

	// Priority Errors
	TaskErrorPriorityEmptyCode
	TaskErrorInvalidPriorityCode
	*/

	// DeviceErrorStatusEmptyCode is a status error
	DeviceErrorStatusEmptyCode
	// DeviceErrorInvalidStatusCode is a status error
	DeviceErrorInvalidStatusCode

	// DeviceErrorDomainEmptyCode gives Domain Errors
	DeviceErrorDomainEmptyCode
	// DeviceErrorInvalidDomainCode gives Domain Errors
	DeviceErrorInvalidDomainCode

	// DeviceErrorCategoryEmptyCode gives Category Errors
	DeviceErrorCategoryEmptyCode
	// DeviceErrorInvalidCategoryCode gives Category Errors
	DeviceErrorInvalidCategoryCode

	// DeviceErrorAssetIDEmptyCode gives Parent UID Errors
	DeviceErrorAssetIDEmptyCode
	// DeviceErrorInvalidAssetIDCode gives Parent UID Errors
	DeviceErrorInvalidAssetIDCode

	// DeviceErrorInventoryIDEmptyCode gives Task Domain Errors
	DeviceErrorInventoryIDEmptyCode
	// DeviceErrorInvalidInventoryIDCode gives Task Domain Errors
	DeviceErrorInvalidInventoryIDCode
	// DeviceErrorInvalidAreaIDCode gives Task Domain Errors
	DeviceErrorInvalidAreaIDCode

	// DeviceErrorDeviceNotFoundCode gives Device General Errors
	DeviceErrorDeviceNotFoundCode
)

// DeviceError is a custom error from Go built-in error
type DeviceError struct {
	Code int
}

func (e DeviceError) Error() string {
	switch e.Code {
	case DeviceErrorTitleEmptyCode:
		return "Device title is required."
	case DeviceErrorIDInvalidCode:
		return "Device ID is invalid."
	case DeviceErrorDescriptionEmptyCode:
		return "Device description is required."
	//case TaskErrorDueDateEmptyCode:
	//	return "Task due date is required."
	//case TaskErrorDueDateInvalidCode:
	//	return "Task due date cannot be earlier than the current date."
	//case TaskErrorPriorityEmptyCode:
	//	return "Task priority is required."
	//case TaskErrorInvalidPriorityCode:
	//	return "Task priority is invalid."
	case DeviceErrorStatusEmptyCode:
		return "Device status is required."
	case DeviceErrorInvalidStatusCode:
		return "Device status is invalid."
	//case DeviceErrorDomainEmptyCode:
	//	return "Device domain is required."
	//case DeviceErrorInvalidDomainCode:
	//	return "Device domain is invalid."
	case DeviceErrorCategoryEmptyCode:
		return "Device category is required."
	case DeviceErrorInvalidCategoryCode:
		return "Device category is invalid."
	case DeviceErrorAssetIDEmptyCode:
		return "Device must have a referenced asset."
	case DeviceErrorInvalidAssetIDCode:
		return "Device asset reference is invalid."
	//case TaskErrorInventoryIDEmptyCode:
	//	return "This Task category requires an inventory reference."
	//case TaskErrorInvalidInventoryIDCode:
	//	return "Task material reference is invalid."
	//case TaskErrorInvalidAreaIDCode:
	//	return " area reference is invalid."
	case DeviceErrorDeviceNotFoundCode:
		return "Device not found"
	default:
		return "Unrecognized Device Error Code"
	}
}