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


// DeviceDomainDetailedCrop interface in domain
// But contains more detailed information of material, area and crop
type DeviceDomainDetailedCrop struct {
	Material *DeviceDomainCropMaterial `json:"material"`
	Area     *DeviceDomainCropArea     `json:"area"`
	Crop     *DeviceDomainCropBatch    `json:"crop"`
}


// DeviceDomainCropArea is used ...
type DeviceDomainCropArea struct {
	AreaID   *uuid.UUID `json:"area_id"`
	AreaName string     `json:"area_name"`
}

// DeviceDomainCropBatch is used...
type DeviceDomainCropBatch struct {
	CropID      *uuid.UUID `json:"crop_id"`
	CropBatchID string     `json:"crop_batch_id"`
}

// DeviceDomainCropMaterial is used ...
type DeviceDomainCropMaterial struct {
	MaterialID           *uuid.UUID `json:"material_id"`
	MaterialName         string     `json:"material_name"`
	MaterialType         string     `json:"material_type"`
	MaterialDetailedType string     `json:"material_detailed_type"`
}

// Code is used ...
func (d DeviceDomainDetailedCrop) Code() string {
	return domain.DeviceDomainCropCode
}

// DeviceDomainDetailedArea is used ...
type DeviceDomainDetailedArea struct {
	MaterialID           *uuid.UUID `json:"material_id"`
	MaterialName         string     `json:"material_name"`
	MaterialType         string     `json:"material_type"`
	MaterialDetailedType string     `json:"material_detailed_type"`
}

// Code is used ...
func (d DeviceDomainDetailedArea) Code() string {
	return domain.DeviceDomainCropCode
}

// DeviceDomainDetailedReservoir is used ...
type DeviceDomainDetailedReservoir struct {
	MaterialID           *uuid.UUID `json:"material_id"`
	MaterialName         string     `json:"material_name"`
	MaterialType         string     `json:"material_type"`
	MaterialDetailedType string     `json:"material_detailed_type"`
}

// Code is used ...
func (d DeviceDomainDetailedReservoir) Code() string {
	return domain.DeviceDomainCropCode
}
