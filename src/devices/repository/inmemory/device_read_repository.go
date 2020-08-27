package inmemory

import (
	"github.com/PalmaPedro/tania-core/src/devices/repository"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
)

// DeviceReadRepositoryInMemory is used ...
type DeviceReadRepositoryInMemory struct {
	Storage *storage.DeviceReadStorage
}

// NewDeviceReadRepositoryInMemory is used ...
func NewDeviceReadRepositoryInMemory(s *storage.DeviceReadStorage) repository.DeviceReadRepository {
	return &DeviceReadRepositoryInMemory{Storage: s}
}

// Save is used ...
func (f *DeviceReadRepositoryInMemory) Save(deviceRead *storage.DeviceRead) <-chan error {
	result := make(chan error)

	go func() {
		f.Storage.Lock.Lock()
		defer f.Storage.Lock.Unlock()

		f.Storage.DeviceReadMap[deviceRead.UID] = *deviceRead

		result <- nil

		close(result)
	}()

	return result
}
