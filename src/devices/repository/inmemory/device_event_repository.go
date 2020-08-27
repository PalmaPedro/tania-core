package inmemory

import (
	"github.com/PalmaPedro/tania-core/src/devices/repository"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
	uuid "github.com/satori/go.uuid"
)

// DeviceEventRepositoryInMemory is used ...
type DeviceEventRepositoryInMemory struct {
	Storage *storage.DeviceEventStorage
}

// NewDeviceEventRepositoryInMemory is used ...
func NewDeviceEventRepositoryInMemory(s *storage.DeviceEventStorage) repository.DeviceEventRepository {
	return &DeviceEventRepositoryInMemory{Storage: s}
}

// Save is to save
func (f *DeviceEventRepositoryInMemory) Save(uid uuid.UUID, latestVersion int, events []interface{}) <-chan error {
	result := make(chan error)

	go func() {
		f.Storage.Lock.Lock()
		defer f.Storage.Lock.Unlock()

		for _, v := range events {
			latestVersion++
			f.Storage.DeviceEvents = append(f.Storage.DeviceEvents, storage.DeviceEvent{
				DeviceUID: uid,
				Version: latestVersion,
				Event:   v,
			})
		}

		result <- nil

		close(result)
	}()

	return result
}
