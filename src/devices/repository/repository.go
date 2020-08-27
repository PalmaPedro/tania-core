package repository

import (
	"github.com/PalmaPedro/tania-core/src/devices/domain"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
	uuid "github.com/satori/go.uuid"
)

// Result is a struct to wrap repository result
// so its easy to use it in channel
type Result struct {
	Result interface{}
	Error  error
}

// EventWrapper is used to wrap the event interface with its struct name,
// so it will be easier to unmarshal later
type EventWrapper struct {
	EventName string
	EventData interface{}
}

// DeviceEventRepository is used ...
type DeviceEventRepository interface {
	Save(uid uuid.UUID, latestVersion int, events []interface{}) <-chan error
}

// BuildDeviceFromEventHistory is used ...
func BuildDeviceFromEventHistory(deviceService domain.DeviceService, events []storage.DeviceEvent) *domain.Device {
	state := &domain.Device{}
	for _, v := range events {
		state.Transition(deviceService, v.Event)
		state.Version++
	}
	return state
}

// DeviceReadRepository is used
type DeviceReadRepository interface {
	Save(deviceRead *storage.DeviceRead) <-chan error
}