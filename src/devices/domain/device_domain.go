package domain

import (
	uuid "github.com/satori/go.uuid"
)

// DeviceDomainAreaCode is used ...
const (
	DeviceDomainAreaCode      = "AREA"
	DeviceDomainCropCode      = "CROP"
	DeviceDomainFinanceCode   = "FINANCE"
	DeviceDomainGeneralCode   = "GENERAL"
	DeviceDomainInventoryCode = "INVENTORY"
	DeviceDomainReservoirCode = "RESERVOIR"
)

// DeviceDomain is used ...
type DeviceDomain interface {
	Code() string
}

// DeviceDomainArea is used ...
type DeviceDomainArea struct {
	MaterialID *uuid.UUID `json:"material_id"`
}

// Code is used ...
func (d DeviceDomainArea) Code() string {
	return DeviceDomainAreaCode
}

// DeviceDomainCrop is used ...
type DeviceDomainCrop struct {
	MaterialID *uuid.UUID `json:"material_id"`
	AreaID     *uuid.UUID `json:"area_id"`
}

// Code is used ...
func (d DeviceDomainCrop) Code() string {
	return DeviceDomainCropCode
}

// DeviceDomainFinance ...
type DeviceDomainFinance struct {
}

// Code is used ...
func (d DeviceDomainFinance) Code() string {
	return DeviceDomainFinanceCode
}

// DeviceDomainGeneral is used ...
type DeviceDomainGeneral struct {
}

// Code is used ...
func (d DeviceDomainGeneral) Code() string {
	return DeviceDomainGeneralCode
}

// DeviceDomainInventory is used ...
type DeviceDomainInventory struct {
}

// Code is used ...
func (d DeviceDomainInventory) Code() string {
	return DeviceDomainInventoryCode
}

// DeviceDomainReservoir is used ...
type DeviceDomainReservoir struct {
	MaterialID *uuid.UUID `json:"material_id"`
}

// Code is used ...
func (d DeviceDomainReservoir) Code() string {
	return DeviceDomainReservoirCode
}

// CreateDeviceDomainArea is used ....
func CreateDeviceDomainArea(deviceService DeviceService, category string, materialID *uuid.UUID) (DeviceDomainArea, error) {

	err := validateDeviceCategory(category)
	if err != nil {
		return DeviceDomainArea{}, err
	}

	if materialID != nil {
		err := validateAssetID(deviceService, materialID, DeviceDomainInventoryCode)
		if err != nil {
			return DeviceDomainArea{}, err
		}
	}

	return DeviceDomainArea{
		MaterialID: materialID,
	}, nil
}

// CreateDeviceDomainCrop is used ...
func CreateDeviceDomainCrop(deviceService DeviceService, category string, materialID *uuid.UUID, areaID *uuid.UUID) (DeviceDomainCrop, error) {

	err := validateDeviceCategory(category)
	if err != nil {
		return DeviceDomainCrop{}, err
	}

	if materialID != nil {
		err := validateAssetID(deviceService, materialID, DeviceDomainInventoryCode)
		if err != nil {
			return DeviceDomainCrop{}, err
		}
	}

	if areaID != nil {
		err := validateAssetID(deviceService, areaID, DeviceDomainAreaCode)
		if err != nil {
			return DeviceDomainCrop{}, err
		}
	}

	return DeviceDomainCrop{
		MaterialID: materialID,
		AreaID:     areaID,
	}, nil
}

// CreateDeviceDomainFinance is used ...
func CreateDeviceDomainFinance() (DeviceDomainFinance, error) {
	return DeviceDomainFinance{}, nil
}

// CreateDeviceDomainGeneral is used ....
func CreateDeviceDomainGeneral() (DeviceDomainGeneral, error) {
	return DeviceDomainGeneral{}, nil
}

// CreateDeviceDomainInventory is used ...
func CreateDeviceDomainInventory() (DeviceDomainInventory, error) {
	return DeviceDomainInventory{}, nil
}

// CreateDeviceDomainReservoir is used ...
func CreateDeviceDomainReservoir(deviceService DeviceService, category string, materialID *uuid.UUID) (DeviceDomainReservoir, error) {

	err := validateDeviceCategory(category)
	if err != nil {
		return DeviceDomainReservoir{}, err
	}

	if materialID != nil {
		err := validateAssetID(deviceService, materialID, DeviceDomainInventoryCode)
		if err != nil {
			return DeviceDomainReservoir{}, err
		}
	}

	return DeviceDomainReservoir{
		MaterialID: materialID,
	}, nil
}

// validateAssetID
func validateDomainAssetID(deviceService DeviceService, assetid *uuid.UUID, devicedomain string) error {
	if devicedomain == "" {
		return DeviceError{DeviceErrorDomainEmptyCode}
	}
	//Find asset in repository
	// if not found return error

	switch devicedomain {
	case DeviceDomainAreaCode:

		serviceResult := deviceService.FindAreaByID(*assetid)

		if serviceResult.Error != nil {
			return serviceResult.Error
		}
	case DeviceDomainCropCode:

		serviceResult := deviceService.FindCropByID(*assetid)

		if serviceResult.Error != nil {
			return serviceResult.Error
		}
	case DeviceDomainInventoryCode:

		serviceResult := deviceService.FindMaterialByID(*assetid)

		if serviceResult.Error != nil {
			return serviceResult.Error
		}
	case DeviceDomainReservoirCode:

		serviceResult := deviceService.FindReservoirByID(*assetid)

		if serviceResult.Error != nil {
			return serviceResult.Error
		}

	default:
		return DeviceError{DeviceErrorInvalidDomainCode}
	}
	return nil
}