package domain

const (
	// DeviceStatusCreated defines status as 'start'
	DeviceStatusCreated = "START_DEVICE"
	// DeviceStatusCancelled defines status as 'stopped'
	DeviceStatusCancelled = "STOP_DEVICE"
	// DeviceStatusCompleted defines status as 'iddle'
	DeviceStatusCompleted = "DEVICE_IDLE"
)

// DeviceStatus is used ...
type DeviceStatus struct {
	Code string
	Name string
}

// FindAllDeviceStatus is used 
func FindAllDeviceStatus() []DeviceStatus {
	return []DeviceStatus{
		DeviceStatus{Code: DeviceStatusCreated, Name: "Start device"},
		DeviceStatus{Code: DeviceStatusCancelled, Name: "Stop device"},
		DeviceStatus{Code: DeviceStatusCompleted, Name: "Device is idle"},
	}
}

// FindDeviceStatusByCode is used ...
func FindDeviceStatusByCode(code string) (DeviceStatus, error) {
	items := FindAllDeviceStatus()

	for _, item := range items {
		if item.Code == code {
			return item, nil
		}
	}

	return DeviceStatus{}, DeviceError{DeviceErrorInvalidStatusCode}
}