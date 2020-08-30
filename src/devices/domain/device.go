package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// DeviceService is used 
type DeviceService interface {
	FindAreaByID(uid uuid.UUID) ServiceResult
	FindCropByID(uid uuid.UUID) ServiceResult
	FindMaterialByID(uid uuid.UUID) ServiceResult
	FindReservoirByID(uid uuid.UUID) ServiceResult
}

// ServiceResult is the container for service result
type ServiceResult struct {
	Result interface{}
	Error  error
}

// Device is used
type Device struct {
	UID           uuid.UUID  `json:"uid"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	CreatedDate   time.Time  `json:"created_date"`
	//DueDate       *time.Time `json:"due_date, omitempty"`
	//CompletedDate *time.Time `json:"completed_date"`
	//CancelledDate *time.Time `json:"cancelled_date"`
	//Priority      string     `json:"priority"`
	Status        string     `json:"status"`
	Domain        string     `json:"domain"`
	DomainDetails DeviceDomain `json:"domain_details"`
	Category      string     `json:"category"`
	//IsDue         bool       `json:"is_due"`
	AssetID       *uuid.UUID `json:"asset_id"`

	// Events
	Version            int
	UncommittedChanges []interface{}
}

// CreateDevice is a function that will add a new device to the system
func CreateDevice(deviceService DeviceService, title string, description string,  duedate *time.Time, devicedomain DeviceDomain, devicecategory string, assetid *uuid.UUID) (*Device, error) {
	// add validation

	err := validateDeviceTitle(title)
	if err != nil {
		return &Device{}, err
	}

	err = validateDeviceDescription(description)
	if err != nil {
		return &Device{}, err
	}

	/*
	err = validateTaskDueDate(duedate)
	if err != nil {
		return &Device{}, err
	}*/

	/*
	err = validateTaskPriority(priority)
	if err != nil {
		return &Device{}, err
	}*/

	err = validateDeviceCategory(devicecategory)
	if err != nil {
		return &Device{}, err
	}

	err = validateAssetID(deviceService, assetid, devicedomain.Code())
	if err != nil {
		return &Device{}, err
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return &Device{}, err
	}

	initial := &Device{}

	initial.TrackChange(deviceService, DeviceCreated{
		Title:         title,
		UID:           uid,
		Description:   description,
		//CreatedDate:   time.Now(),
		//DueDate:       duedate,
		//Priority:      priority,
		Status:        DeviceStatusCreated,
		Domain:        devicedomain.Code(),
		DomainDetails: devicedomain,
		Category:      devicecategory,
		//IsDue:         false,
		AssetID:       assetid,
	})

	return initial, nil
}

// ChangeDeviceTitle is a function that allows the user to change title of a device
func (t *Device) ChangeDeviceTitle(deviceService DeviceService, title string) (*Device, error) {
	err := validateDeviceTitle(title)
	if err != nil {
		return &Device{}, err
	}

	event := DeviceTitleChanged{
		UID:   t.UID,
		Title: title,
	}

	t.TrackChange(deviceService, event)

	return t, nil
}

// ChangeDeviceDescription is a function used to update a device's description
func (t *Device) ChangeDeviceDescription(deviceService DeviceService, description string) (*Device, error) {
	err := validateDeviceDescription(description)
	if err != nil {
		return &Device{}, err
	}

	event := DeviceDescriptionChanged{
		UID:         t.UID,
		Description: description,
	}

	t.TrackChange(deviceService, event)

	return t, nil
}

// ChangeDeviceCategory is a function used to update a device's category
func (t *Device) ChangeDeviceCategory(deviceService DeviceService, category string) (*Device, error) {
	err := validateDeviceCategory(category)
	if err != nil {
		return &Device{}, err
	}

	event := DeviceCategoryChanged{
		UID:      t.UID,
		Category: category,
	}

	t.TrackChange(deviceService, event)

	return t, nil
}

// ChangeDeviceDetails is a function used to update a device's details
func (t *Device) ChangeDeviceDetails(deviceService DeviceService, details DeviceDomain) (*Device, error) {

	event := DeviceDetailsChanged{
		UID:           t.UID,
		DomainDetails: details,
	}

	t.TrackChange(deviceService, event)

	return t, nil
}


// TrackChange is used 
func (t *Device) TrackChange(deviceService DeviceService, event interface{}) error {
	t.UncommittedChanges = append(t.UncommittedChanges, event)
	err := t.Transition(deviceService, event)
	if err != nil {
		return err
	}

	return nil
}

// Transition is used
func (t *Device) Transition(deviceService DeviceService, event interface{}) error {
	switch e := event.(type) {
	case DeviceCreated:
		t.Title = e.Title
		t.UID = e.UID
		t.Description = e.Description
		//state.CreatedDate = e.CreatedDate
		//state.DueDate = e.DueDate
		//state.Priority = e.Priority
		t.Status = e.Status
		t.Domain = e.Domain
		t.DomainDetails = e.DomainDetails
		t.Category = e.Category
		//state.IsDue = e.IsDue
		t.AssetID = e.AssetID
	case DeviceTitleChanged:
		t.Title = e.Title
	case DeviceDescriptionChanged:
		t.Description = e.Description
	//case TaskDueDateChanged:
	//	state.DueDate = e.DueDate
	//case TaskPriorityChanged:
	//	state.Priority = e.Priority
	//case TaskCategoryChanged:
	//	state.Category = e.Category
	case DeviceDetailsChanged:
		t.DomainDetails = e.DomainDetails
	//case TaskCancelled:
	//	state.CancelledDate = e.CancelledDate
	//	state.Status = TaskStatusCancelled
	//case TaskCompleted:
	//	state.CompletedDate = e.CompletedDate
	//	state.Status = TaskStatusCompleted
	//case TaskDue:
	//	state.IsDue = true
	}

	return nil
}

// validateDeviceTitle
func validateDeviceTitle(title string) error {
	if title == "" {
		return DeviceError{DeviceErrorTitleEmptyCode}
	}
	return nil
}

// validateTaskDescription
func validateDeviceDescription(description string) error {
	if description == "" {
		return DeviceError{DeviceErrorDescriptionEmptyCode}
	}
	return nil
}

/*
// validateDeviceDueDate
func validateDeviceDueDate(newdate *time.Time) error {
	if newdate != nil {
		if (*newdate).Before(time.Now()) {
			return DeviceError{DeviceErrorDueDateInvalidCode}
		}
	}
	return nil
}*/

/*
//validateTaskPriority
func validateTaskPriority(priority string) error {

	if priority == "" {
		return TaskError{TaskErrorPriorityEmptyCode}
	}

	_, err := FindTaskPriorityByCode(priority)
	if err != nil {
		return err
	}

	return nil
}*/

// validateDeviceStatus
func validateDeviceStatus(status string) error {

	if status == "" {
		return DeviceError{DeviceErrorStatusEmptyCode}
	}

	_, err := FindDeviceStatusByCode(status)
	if err != nil {
		return err
	}

	return nil
}

// validateDeviceCategory
func validateDeviceCategory(devicecategory string) error {

	if devicecategory == "" {
		return DeviceError{DeviceErrorCategoryEmptyCode}
	}

	_, err := FindDeviceCategoryByCode(devicecategory)
	if err != nil {
		return err
	}

	return nil
}

// validateAssetID
func validateAssetID(deviceService DeviceService, assetid *uuid.UUID, devicedomain string) error {

	if assetid != nil {
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
		/*
		case DeviceDomainReservoirCode:
			serviceResult := deviceService.FindReservoirByID(*assetid)

			if serviceResult.Error != nil {
				return serviceResult.Error
			}*/
		default:
			return DeviceError{DeviceErrorInvalidDomainCode}
		}
	}
	return nil
}






