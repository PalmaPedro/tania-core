package inmemory

import (
	"sort"

	"github.com/PalmaPedro/tania-core/src/devices/query"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
	uuid "github.com/satori/go.uuid"
)

// DeviceEventQueryInMemory is used ...
type DeviceEventQueryInMemory struct {
	Storage *storage.DeviceEventStorage
}

// NewDeviceEventQueryInMemory is used ...
func NewDeviceEventQueryInMemory(s *storage.DeviceEventStorage) query.DeviceEventQuery {
	return &DeviceEventQueryInMemory{Storage: s}
}

// FindAllByDeviceID is used ...
func (f *DeviceEventQueryInMemory) FindAllByDeviceID(uid uuid.UUID) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		f.Storage.Lock.RLock()
		defer f.Storage.Lock.RUnlock()

		events := []storage.DeviceEvent{}
		for _, v := range f.Storage.DeviceEvents {
			if v.DeviceUID == uid {
				events = append(events, v)
			}
		}

		sort.Slice(events, func(i, j int) bool {
			return events[i].Version < events[j].Version
		})

		result <- query.Result{Result: events}
	}()

	return result
}