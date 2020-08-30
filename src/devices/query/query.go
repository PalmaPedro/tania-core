package query

import (
	//assetsdomain "github.com/PalmaPedro/tania-core/src/assets/domain"
	uuid "github.com/satori/go.uuid"
)

// Result is used...
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

// AreaQuery is used ...
type AreaQuery interface {
	FindByID(areaUID uuid.UUID) <-chan Result
}

// CropQuery is used ...
type CropQuery interface {
	FindCropByID(cropUID uuid.UUID) <-chan Result
}

// MaterialQuery is used ...
type MaterialQuery interface {
	FindMaterialByID(materialID uuid.UUID) <-chan Result
}


// DeviceEventQuery is used ...
type DeviceEventQuery interface {
	FindAllByDeviceID(uid uuid.UUID) <-chan Result
}

// DeviceReadQuery is used ...
type DeviceReadQuery interface {
	FindAll(page, limit int) <-chan Result
	FindByID(taskUID uuid.UUID) <-chan Result
	FindDevicesWithFilter(params map[string]string, page, limit int) <-chan Result
  CountAll() <-chan Result
  CountDevicesWithFilter(params map[string]string) <-chan Result
}

// DeviceQuery is used ...
type DeviceQuery interface {
	FindDeviceByID(deviceUID uuid.UUID) <-chan Result
}

// ReservoirQuery is used ...
type ReservoirQuery interface {
	FindReservoirByID(reservoirUID uuid.UUID) <-chan Result
}

/*
TODO

type DeviceQuery interface {
	FindDeviceByID(deviceUID uuid.UUID) <-chan QueryResult
}

type FinanceQuery interface {
	FindFinanceByID(financeUID uuid.UUID) <-chan QueryResult
}

*/

// QUERY RESULTS

// DeviceAreaQueryResult is used
type DeviceAreaQueryResult struct {
	UID  uuid.UUID `json:"uid"`
	Name string    `json:"name"`
}

// DeviceCropQueryResult is used
type DeviceCropQueryResult struct {
	UID     uuid.UUID `json:"uid"`
	BatchID string    `json:"batch_id"`
}

// DeviceMaterialQueryResult is used 
type DeviceMaterialQueryResult struct {
	UID              uuid.UUID `json:"uid"`
	TypeCode         string    `json:"type"`
	DetailedTypeCode string    `json:"detailed_type"`
	Name             string    `json:"name"`
}

// DeviceReservoirQueryResult is used
type DeviceReservoirQueryResult struct {
	UID  uuid.UUID `json:"uid"`
	Name string    `json:"name"`
}