package server

import (
	"errors"
	"net/http"

	"github.com/PalmaPedro/tania-core/src/devices/domain"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// SaveToDeviceReadModel is used ...
func (s *DeviceServer) SaveToDeviceReadModel(event interface{}) error {
	deviceRead := &storage.DeviceRead{}

	switch e := event.(type) {
	case domain.DeviceCreated:

		deviceRead.Title = e.Title
		deviceRead.UID = e.UID
		deviceRead.Description = e.Description
		deviceRead.CreatedDate = e.CreatedDate
		//taskRead.DueDate = e.DueDate
		//taskRead.Priority = e.Priority
		deviceRead.Status = e.Status
		deviceRead.Domain = e.Domain
		deviceRead.DomainDetails = e.DomainDetails
		deviceRead.Category = e.Category
		//taskRead.IsDue = e.IsDue
		deviceRead.AssetID = e.AssetID
	case domain.DeviceTitleChanged:

		// Get DeviceRead By UID
		deviceReadFromRepo, err := s.getDeviceReadFromID(e.UID)
		if err != nil {
			return err
		}

		deviceReadFromRepo.Title = e.Title
		deviceRead = deviceReadFromRepo
	case domain.DeviceDescriptionChanged:

		// Get DeviceRead By UID
		deviceReadFromRepo, err := s.getDeviceReadFromID(e.UID)
		if err != nil {
			return err
		}

		deviceReadFromRepo.Description = e.Description
		deviceRead = deviceReadFromRepo
	/*	
	case domain.DevicePriorityChanged:

		// Get DeviceRead By UID
		deviceReadFromRepo, err := s.getDeviceReadFromID(e.UID)
		if err != nil {
			return err
		}

		deviceReadFromRepo.Priority = e.Priority
		deviceRead = deviceReadFromRepo
	case domain.DeviceDueDateChanged:

		// Get DeviceRead By UID
		deviceReadFromRepo, err := s.getDeviceReadFromID(e.UID)
		if err != nil {
			return err
		}

		//deviceReadFromRepo.DueDate = e.DueDate
		deviceRead = deviceReadFromRepo
	*/
	case domain.DeviceCategoryChanged:

		// Get DeviceRead By UID
		deviceReadFromRepo, err := s.getDeviceReadFromID(e.UID)
		if err != nil {
			return err
		}

		deviceReadFromRepo.Category = e.Category
		deviceRead = deviceReadFromRepo
	case domain.DeviceDetailsChanged:

		// Get DeviceRead By UID
		deviceReadFromRepo, err := s.getDeviceReadFromID(e.UID)
		if err != nil {
			return err
		}

		deviceReadFromRepo.DomainDetails = e.DomainDetails
		deviceRead = deviceReadFromRepo
	
	/*
	case domain.DeviceCompleted:

		// Get DeviceRead By UID
		deviceReadFromRepo, err := s.getDeviceReadFromID(e.UID)
		if err != nil {
			return err
		}

		deviceReadFromRepo.CompletedDate = e.CompletedDate
		deviceReadFromRepo.Status = domain.DeviceStatusCompleted
		deviceRead = deviceReadFromRepo

	case domain.DeviceCancelled:

		// Get DeviceRead By UID

		deviceReadFromRepo, err := s.getDeviceReadFromID(e.UID)
		if err != nil {
			return err
		}

		deviceReadFromRepo.CancelledDate = e.CancelledDate
		deviceReadFromRepo.Status = domain.DeviceStatusCancelled
		deviceRead = deviceReadFromRepo

	case domain.TaskDue:

		// Get TaskRead By UID

		taskReadFromRepo, err := s.getTaskReadFromID(e.UID)
		if err != nil {
			return err
		}

		taskReadFromRepo.IsDue = true
		taskRead = taskReadFromRepo
		*/
	default:
		return errors.New("Unknown device event")
	}

	err := <-s.DeviceReadRepo.Save(deviceRead)
	if err != nil {
		return err
	}

	return nil
}

func (s *DeviceServer) getDeviceReadFromID(uid uuid.UUID) (*storage.DeviceRead, error) {

	readResult := <-s.DeviceReadQuery.FindByID(uid)

	deviceReadFromRepo, ok := readResult.Result.(storage.DeviceRead)

	if deviceReadFromRepo.UID != uid {
		return &storage.DeviceRead{}, domain.DeviceError{domain.DeviceErrorDeviceNotFoundCode}
	}
	if !ok {
		return &storage.DeviceRead{}, echo.NewHTTPError(http.StatusBadRequest, "Internal server error")
	} 
	return &deviceReadFromRepo, nil
}