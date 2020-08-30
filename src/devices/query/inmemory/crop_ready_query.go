package inmemory

import (
	"github.com/Tanibox/tania-core/src/growth/storage"
	"github.com/PalmaPedro/tania-core/src/devices/query"
	uuid "github.com/satori/go.uuid"
)

// CropQueryInMemory is used ...
type CropQueryInMemory struct {
	Storage *storage.CropReadStorage
}

// NewCropQueryInMemory is used ...
func NewCropQueryInMemory(s *storage.CropReadStorage) query.CropQuery {
	return CropQueryInMemory{Storage: s}
}

// FindCropByID is used ...
func (s CropQueryInMemory) FindCropByID(uid uuid.UUID) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		s.Storage.Lock.RLock()
		defer s.Storage.Lock.RUnlock()

		crop := query.DeviceCropQueryResult{}

		for _, val := range s.Storage.CropReadMap {
			if val.UID == uid {
				crop.UID = uid
				crop.BatchID = val.BatchID
			}
		}
		result <- query.Result{Result: crop}

		close(result)
	}()

	return result
}