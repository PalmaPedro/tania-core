package domain

import (

	//"time"

	uuid "github.com/satori/go.uuid"
	
)

// DeviceCreatedCode is used
const (
	DeviceCreatedCode            = "DeviceCreated"
	DeviceTitleChangedCode       = "DeviceTitleChanged"
	DeviceDescriptionChangedCode = "DeviceDescriptionChanged"
	//DevicePriorityChangedCode    = "TaskPriorityChanged"
	//DeviceDueDateChangedCode     = "TaskDueDateChanged"
	DeviceCategoryChangedCode    = "DeviceCategoryChanged"
	DeviceDetailsChangedCode     = "DeviceDetailsChanged"
	DeviceAssetIDChangedCode     = "DeviceAssetIDChanged"
	//DeviceCompletedCode          = "DeviceCompleted"
	//DeviceCancelledCode          = "DeviceCancelled"
	//DeviceDueCode                = "DeviceDue"
)

// DeviceCreated is used
type DeviceCreated struct {
	UID           uuid.UUID  `json:"uid"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	//CreatedDate   time.Time  `json:"created_date"`
	//DueDate       *time.Time `json:"due_date"`
	//Priority      string     `json:"priority"`
	Status        string     `json:"status"`
	Domain        string     `json:"domain"`
	DomainDetails DeviceDomain `json:"domain_details"`
	Category      string     `json:"category"`
	//IsDue         bool       `json:"is_due"`
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

/*
type DevicePriorityChanged struct {
	UID      uuid.UUID `json:"uid"`
	Priority string    `json:"priority"`
}*/

/*
type TaskDueDateChanged struct {
	UID     uuid.UUID  `json:"uid"`
	DueDate *time.Time `json:"due_date"`
}*/

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

/*
type TaskCompleted struct {
	UID           uuid.UUID  `json:"uid"`
	Status        string     `json:"status"`
	CompletedDate *time.Time `json:"completed_date"`
}

type TaskCancelled struct {
	UID           uuid.UUID  `json:"uid"`
	Status        string     `json:"status"`
	CancelledDate *time.Time `json:"cancelled_date"`
}*/

/*
type TaskDue struct {
	UID uuid.UUID `json:"uid"`
}*/