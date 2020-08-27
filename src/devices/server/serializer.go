package server

import (
	"github.com/PalmaPedro/tania-core/src/devices/domain"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
)

// MapDeviceToDeviceRead is used ...
func MapDeviceToDeviceRead(device *domain.Device) *storage.DeviceRead {
	deviceRead := &storage.DeviceRead{
		Title:         device.Title,
		UID:           device.UID,
		Description:   device.Description,
		CreatedDate:   device.CreatedDate,
		//DueDate:       task.DueDate,
		//CompletedDate: task.CompletedDate,
		//CancelledDate: task.CancelledDate,
		//Priority:      device.Priority,
		Status:        device.Status,
		Domain:        device.Domain,
		DomainDetails: device.DomainDetails,
		Category:      device.Category,
		//IsDue:         task.IsDue,
		AssetID:       device.AssetID,
	}
	return deviceRead
}