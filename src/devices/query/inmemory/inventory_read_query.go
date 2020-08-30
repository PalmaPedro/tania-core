package inmemory

import (
	assetsdomain "github.com/Tanibox/tania-core/src/assets/domain"
	assetsstorage "github.com/Tanibox/tania-core/src/assets/storage"
	"github.com/PalmaPedro/tania-core/src/devices/query"
	uuid "github.com/satori/go.uuid"
)

// MaterialQueryInMemory is used ...
type MaterialQueryInMemory struct {
	Storage *assetsstorage.MaterialReadStorage
}

// NewMaterialQueryInMemory is used ...
func NewMaterialQueryInMemory(s *assetsstorage.MaterialReadStorage) query.MaterialQuery {
	return MaterialQueryInMemory{Storage: s}
}

// FindMaterialByID is used ...
func (s MaterialQueryInMemory) FindMaterialByID(inventoryUID uuid.UUID) <-chan query.Result {
	result := make(chan query.Result)
	go func() {
		s.Storage.Lock.RLock()
		defer s.Storage.Lock.RUnlock()

		ci := query.DeviceMaterialQueryResult{}
		for _, val := range s.Storage.MaterialReadMap {
			// WARNING, domain leakage

			if val.UID == inventoryUID {
				ci.UID = val.UID
				ci.Name = val.Name
				ci.TypeCode = val.Type.Code()

				switch v := val.Type.(type) {
				case assetsdomain.MaterialTypeSeed:
					ci.DetailedTypeCode = v.PlantType.Code
				case assetsdomain.MaterialTypePlant:
					ci.DetailedTypeCode = v.PlantType.Code
				case assetsdomain.MaterialTypeAgrochemical:
					ci.DetailedTypeCode = v.ChemicalType.Code
				}
			}
		}
		result <- query.Result{Result: ci}

		close(result)
	}()

	return result
}