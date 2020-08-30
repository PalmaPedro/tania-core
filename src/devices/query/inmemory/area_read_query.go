package inmemory

import (
	"github.com/Tanibox/tania-core/src/assets/storage"
	"github.com/PalmaPedro/tania-core/src/devices/query"
	uuid "github.com/satori/go.uuid"
)

// AreaQueryInMemory is used ...
type AreaQueryInMemory struct {
	Storage *storage.AreaReadStorage
}

// NewAreaQueryInMemory is used
func NewAreaQueryInMemory(s *storage.AreaReadStorage) query.AreaQuery {
	return AreaQueryInMemory{Storage: s}
}

// FindByID is used ...
func (s AreaQueryInMemory) FindByID(uid uuid.UUID) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		s.Storage.Lock.RLock()
		defer s.Storage.Lock.RUnlock()

		area := query.DeviceAreaQueryResult{}
		for _, val := range s.Storage.AreaReadMap {
			if val.UID == uid {
				area.UID = uid
				area.Name = val.Name
			}
		}

		result <- query.Result{Result: area}

		close(result)
	}()

	return result
}