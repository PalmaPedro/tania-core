package service

import (
	domain "github.com/PalmaPedro/tania-core/src/devices/domain"
	"github.com/PalmaPedro/tania-core/src/devices/query"
	uuid "github.com/satori/go.uuid"
)

// DeviceServiceSqLLite handles task behaviours that needs external interaction to be worked
type DeviceServiceSqLLite struct {
	CropQuery      query.CropQuery
	AreaQuery      query.AreaQuery
	MaterialQuery  query.MaterialQuery
	//ReservoirQuery query.ReservoirQuery
}

// FindAreaByID used to query the database and return a result if any
func (s DeviceServiceSqLLite) FindAreaByID(uid uuid.UUID) domain.ServiceResult {
	result := <-s.AreaQuery.FindByID(uid)

	if result.Error != nil {
		return domain.ServiceResult{
			Error: result.Error,
		}
	}

	area, ok := result.Result.(query.DeviceAreaQueryResult)
	if !ok {
		return domain.ServiceResult{
			Error: domain.DeviceError{Code: domain.DeviceErrorInvalidAssetIDCode},
		}
	}

	if area == (query.DeviceAreaQueryResult{}) {
		return domain.ServiceResult{
			Error: domain.DeviceError{Code: domain.DeviceErrorInvalidAssetIDCode},
		}
	}

	return domain.ServiceResult{
		Result: area,
	}
}

// FindCropByID is used ...
func (s DeviceServiceSqLLite) FindCropByID(uid uuid.UUID) domain.ServiceResult {
	result := <-s.CropQuery.FindCropByID(uid)

	if result.Error != nil {
		return domain.ServiceResult{
			Error: result.Error,
		}
	}

	crop, ok := result.Result.(query.DeviceCropQueryResult)
	if !ok {
		return domain.ServiceResult{
			Error: domain.DeviceError{Code: domain.DeviceErrorInvalidAssetIDCode},
		}
	}

	if crop == (query.DeviceCropQueryResult{}) {
		return domain.ServiceResult{
			Error: domain.DeviceError{Code: domain.DeviceErrorInvalidAssetIDCode},
		}
	}

	return domain.ServiceResult{
		Result: crop,
	}
}

// FindMaterialByID is used ..
func (s DeviceServiceSqLLite) FindMaterialByID(uid uuid.UUID) domain.ServiceResult {
	result := <-s.MaterialQuery.FindMaterialByID(uid)

	if result.Error != nil {
		return domain.ServiceResult{
			Error: result.Error,
		}
	}

	material, ok := result.Result.(query.DeviceMaterialQueryResult)
	if !ok {
		return domain.ServiceResult{
			Error: domain.DeviceError{Code: domain.DeviceErrorInvalidAssetIDCode},
		}
	}

	if material == (query.DeviceMaterialQueryResult{}) {
		return domain.ServiceResult{
			Error: domain.DeviceError{Code: domain.DeviceErrorInvalidAssetIDCode},
		}
	}

	return domain.ServiceResult{
		Result: material,
	}
}


/*
func (s DeviceServiceSqLLite) FindReservoirByID(uid uuid.UUID) domain.ServiceResult {
	result := <-s.ReservoirQuery.FindReservoirByID(uid)

	if result.Error != nil {
		return domain.ServiceResult{
			Error: result.Error,
		}
	}

	reservoir, ok := result.Result.(query.DeviceReservoirQueryResult)
	if !ok {
		return domain.ServiceResult{
			Error: domain.DeviceError{Code: domain.DeviceErrorInvalidAssetIDCode},
		}
	}

	if reservoir == (query.DeviceReservoirQueryResult{}) {
		return domain.ServiceResult{
			Error: domain.DeviceError{Code: domain.DeviceErrorInvalidAssetIDCode},
		}
	}

	return domain.ServiceResult{
		Result: reservoir,
	}
}*/