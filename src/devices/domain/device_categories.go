package domain

const (
	// DeviceCategoryRobot is a variable for category of a device
	DeviceCategoryRobot = "ROBOT"
	// DeviceCategorySensor is a variable for a category of a device
	DeviceCategorySensor = "SENSOR"
)

// DeviceCategory defines ...
type DeviceCategory struct {
	Code string
	Name string
}

// FindAllDeviceCategories allows a user to filter a search 
func FindAllDeviceCategories() []DeviceCategory {
	return []DeviceCategory{
		DeviceCategory{Code: DeviceCategoryRobot, Name: "Robot"},
		DeviceCategory{Code: DeviceCategorySensor, Name: "Sensor"},
		
	}
}

// FindDeviceCategoryByCode allows a user to filter a search
func FindDeviceCategoryByCode(code string) (DeviceCategory, error) {
	items := FindAllDeviceCategories()

	for _, item := range items {
		if item.Code == code {
			return item, nil
		}
	}

	return DeviceCategory{}, DeviceError{DeviceErrorInvalidCategoryCode}
}