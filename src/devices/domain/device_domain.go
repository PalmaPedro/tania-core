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
	//DeviceDomainReservoirCode = "RESERVOIR"
)

// DeviceDomain is used in
type DeviceDomain interface {
	Code() string
}

// DeviceDomainArea is used ...
type DeviceDomainArea struct {
	MaterialID *uuid.UUID `json:"material_id"`
}

// Code is used for
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

/*
// FINANCE
type TaskDomainFinance struct {
}
*/

/*
func (d TaskDomainFinance) Code() string {
	return TaskDomainFinanceCode
}
*/


// DeviceDomainGeneral is used
type DeviceDomainGeneral struct {
}

// Code is used ...
func (d DeviceDomainGeneral) Code() string {
	return DeviceDomainGeneralCode
}

// DeviceDomainInventory is used
type DeviceDomainInventory struct {
}

// Code ...
func (d DeviceDomainInventory) Code() string {
	return DeviceDomainInventoryCode
}

/*
// RESERVOIR
type DeviceDomainReservoir struct {
	MaterialID *uuid.UUID `json:"material_id"`
}
*/

/*
// Code ..
func (d DeviceDomainReservoir) Code() string {
	return DeviceDomainReservoirCode
}
*/

// CreateDeviceDomainArea is used
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