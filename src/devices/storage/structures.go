package storage

import (
	"time"

	domain "github.com/PalmaPedro/tania-core/src/devices/domain"
	uuid "github.com/satori/go.uuid"
)

// DeviceEvent is used ...
type DeviceEvent struct {
	DeviceUID     uuid.UUID
	Version     int
	CreatedDate time.Time
	Event       interface{}
}

// DeviceRead is used ...
type DeviceRead struct {
	Title         string            `json:"title"`
	UID           uuid.UUID         `json:"uid"`
	Description   string            `json:"description"`
	CreatedDate   time.Time         `json:"created_date"`
	//DueDate       *time.Time        `json:"due_date, omitempty"`
	//CompletedDate *time.Time        `json:"completed_date"`
	//CancelledDate *time.Time        `json:"cancelled_date"`
	//Priority      string            `json:"priority"`
	Status        string            `json:"status"`
	Domain        string            `json:"domain"`
	DomainDetails domain.DeviceDomain `json:"domain_details"`
	Category      string            `json:"category"`
	//IsDue         bool              `json:"is_due"`
	AssetID       *uuid.UUID        `json:"asset_id"`
}

/*
// Implements DeviceDomainDetailedCrop interface in domain
// But contains more detailed information of material, area and crop
type DeviceDomainDetailedCrop struct {
	Material *TaskDomainCropMaterial `json:"material"`
	Area     *TaskDomainCropArea     `json:"area"`
	Crop     *TaskDomainCropBatch    `json:"crop"`
}*/

/*
type TaskDomainCropArea struct {
	AreaID   *uuid.UUID `json:"area_id"`
	AreaName string     `json:"area_name"`
}

type TaskDomainCropBatch struct {
	CropID      *uuid.UUID `json:"crop_id"`
	CropBatchID string     `json:"crop_batch_id"`
}

type TaskDomainCropMaterial struct {
	MaterialID           *uuid.UUID `json:"material_id"`
	MaterialName         string     `json:"material_name"`
	MaterialType         string     `json:"material_type"`
	MaterialDetailedType string     `json:"material_detailed_type"`
}

func (d TaskDomainDetailedCrop) Code() string {
	return domain.TaskDomainCropCode
}

type TaskDomainDetailedArea struct {
	MaterialID           *uuid.UUID `json:"material_id"`
	MaterialName         string     `json:"material_name"`
	MaterialType         string     `json:"material_type"`
	MaterialDetailedType string     `json:"material_detailed_type"`
}

func (d TaskDomainDetailedArea) Code() string {
	return domain.TaskDomainCropCode
}

type TaskDomainDetailedReservoir struct {
	MaterialID           *uuid.UUID `json:"material_id"`
	MaterialName         string     `json:"material_name"`
	MaterialType         string     `json:"material_type"`
	MaterialDetailedType string     `json:"material_detailed_type"`
}

func (d TaskDomainDetailedReservoir) Code() string {
	return domain.TaskDomainCropCode
}
*/