package domain

import (

	"time"
	uuid "github.com/satori/go.uuid"
	
)

// DeviceCreatedCode is used
const (
	DeviceCreatedCode            = "DeviceCreated"
	DeviceTitleChangedCode       = "DeviceTitleChanged"
	DeviceDescriptionChangedCode = "DeviceDescriptionChanged"
	DeviceCategoryChangedCode    = "DeviceCategoryChanged"
	DeviceDetailsChangedCode     = "DeviceDetailsChanged"
	DeviceAssetIDChangedCode     = "DeviceAssetIDChanged"
)

// DeviceCreated is used
type DeviceCreated struct {
	UID           uuid.UUID  `json:"uid"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	CreatedDate   time.Time  `json:"created_date"`
	Status        string     `json:"status"`
	Domain        string     `json:"domain"`
	DomainDetails DeviceDomain `json:"domain_details"`
	Category      string     `json:"category"`
	AssetID       *uuid.UUID `json:"asset_id"`
}

// DeviceTitleChanged is used ...
type DeviceTitleChanged struct {
	UID   uuid.UUID `json:"uid"`
	Title string    `json:"title"`
}

// DeviceDescriptionChanged is used ...
type DeviceDescriptionChanged struct {
	UID         uuid.UUID `json:"uid"`
	Description string    `json:"description"`
}

// DeviceCategoryChanged is used
type DeviceCategoryChanged struct {
	UID      uuid.UUID `json:"uid"`
	Category string    `json:"category"`
}

// DeviceDetailsChanged is used
type DeviceDetailsChanged struct {
	UID           uuid.UUID  `json:"uid"`
	DomainDetails DeviceDomain `json:"domain_details"`
}

// DeviceAssetIDChanged is used
type DeviceAssetIDChanged struct {
	UID     uuid.UUID  `json:"uid"`
	AssetID *uuid.UUID `json:"asset_id"`
}
