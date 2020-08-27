package decoder

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/PalmaPedro/tania-core/src/devices/domain"
	"github.com/mitchellh/mapstructure"
	uuid "github.com/satori/go.uuid"
)

// DeviceEventWrapper is used...
type DeviceEventWrapper InterfaceWrapper

// UnmarshalJSON is used ...
func (w *DeviceEventWrapper) UnmarshalJSON(b []byte) error {
	wrapper := InterfaceWrapper{}

	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}

	mapped := wrapper.Data.(map[string]interface{})

	f := mapstructure.ComposeDecodeHookFunc(
		UIDHook(),
		TimeHook(time.RFC3339),
		DeviceDomainDetailHook(),
	)

	switch wrapper.Name {
	case domain.DeviceCreatedCode:
		e := domain.DeviceCreated{}

		_, err := Decode(f, &mapped, &e)
		if err != nil {
			return err
		}

		if v, ok := mapped["domain_details"]; ok {
			domainCode, ok := mapped["domain"].(string)
			if !ok {
				return errors.New("Error type assertion")
			}

			e.DomainDetails, err = makeDomainDetails(v, domainCode)
			if err != nil {
				return err
			}
		}

		w.Data = e

	case domain.DeviceTitleChangedCode:
		e := domain.DeviceTitleChanged{}

		_, err := Decode(f, &mapped, &e)
		if err != nil {
			return err
		}

		w.Data = e

	case domain.DeviceDescriptionChangedCode:
		e := domain.DeviceDescriptionChanged{}

		_, err := Decode(f, &mapped, &e)
		if err != nil {
			return err
		}

		w.Data = e


	//case domain.TaskPriorityChangedCode:
	//	e := domain.TaskPriorityChanged{}

	//	_, err := Decode(f, &mapped, &e)
	//	if err != nil {
	//		return err
	//	}

	//	w.Data = e

	//case domain.TaskDueDateChangedCode:
	//	e := domain.TaskDueDateChanged{}

	//	_, err := Decode(f, &mapped, &e)
	//	if err != nil {
	//		return err
	//	}

		w.Data = e


	case domain.DeviceCategoryChangedCode:
		e := domain.DeviceCategoryChanged{}

		_, err := Decode(f, &mapped, &e)
		if err != nil {
			return err
		}

		w.Data = e


	case domain.DeviceDetailsChangedCode:
		e := domain.DeviceDetailsChanged{}

		_, err := Decode(f, &mapped, &e)
		if err != nil {
			return err
		}

		if v, ok := mapped["domain_details"]; ok {
			domainCode, ok := mapped["domain"].(string)
			if !ok {
				return errors.New("Error type assertion")
			}

			e.DomainDetails, err = makeDomainDetails(v, domainCode)
			if err != nil {
				return err
			}
		}

		w.Data = e

	case domain.DeviceAssetIDChangedCode:
		e := domain.DeviceAssetIDChanged{}

		_, err := Decode(f, &mapped, &e)
		if err != nil {
			return err
		}

		w.Data = e

	}

	return nil
}


/*
	case domain.DeviceCompletedCode:
		e := domain.TaskCompleted{}

		_, err := Decode(f, &mapped, &e)
		if err != nil {
			return err
		}

		w.Data = e

	case domain.TaskCancelledCode:
		e := domain.TaskCancelled{}

		_, err := Decode(f, &mapped, &e)
		if err != nil {
			return err
		}

		w.Data = e

	case domain.TaskDueCode:
		e := domain.TaskDue{}

		_, err := Decode(f, &mapped, &e)
		if err != nil {
			return err
		}

		w.Data = e

	}
*/

func makeDomainDetails(v interface{}, domainCode string) (domain.DeviceDomain, error) {
	mapped := v.(map[string]interface{})

	var domainDetails domain.DeviceDomain
	switch domainCode {
	case domain.DeviceDomainAreaCode:
		deviceDomainArea := domain.DeviceDomainArea{}

		if v2, ok2 := mapped["material_id"]; ok2 {
			val, ok2 := v2.(string)
			if !ok2 {
				return domain.DeviceDomainArea{}, nil
			}

			uid, err := uuid.FromString(val)
			if err != nil {
				return domain.DeviceDomainArea{}, err
			}

			deviceDomainArea.MaterialID = &uid
		}

		domainDetails = deviceDomainArea

	case domain.DeviceDomainCropCode:
		deviceDomainCrop := domain.DeviceDomainCrop{}

		if v2, ok2 := mapped["material_id"]; ok2 {
			val, ok2 := v2.(string)
			if !ok2 {
				return domain.DeviceDomainCrop{}, nil
			}

			uid, err := uuid.FromString(val)
			if err != nil {
				return domain.DeviceDomainCrop{}, err
			}

			deviceDomainCrop.MaterialID = &uid
		}
		if v2, ok2 := mapped["area_id"]; ok2 {
			val, ok := v2.(string)
			if !ok {
				return domain.DeviceDomainCrop{}, nil
			}

			uid, err := uuid.FromString(val)
			if err != nil {
				return domain.DeviceDomainCrop{}, err
			}

			deviceDomainCrop.AreaID = &uid
		}

		domainDetails = deviceDomainCrop
	//case domain.DeviceDomainFinanceCode:
	//	domainDetails = domain.DeviceDomainFinance{}
	case domain.DeviceDomainGeneralCode:
		domainDetails = domain.DeviceDomainGeneral{}
	//case domain.TaskDomainInventoryCode:
	//	domainDetails = domain.TaskDomainInventory{}
	//case domain.TaskDomainReservoirCode:
	//	taskDomainReservoir := domain.TaskDomainReservoir{}


	//	if v2, ok2 := mapped["material_id"]; ok2 {
	//		val, ok2 := v2.(string)
	//		if !ok2 {
	//			return domain.TaskDomainReservoir{}, nil
	//		}

	//		uid, err := uuid.FromString(val)
	//		if err != nil {
	//			return domain.TaskDomainReservoir{}, err
	//		}

	//		taskDomainReservoir.MaterialID = &uid
	//	}

	//	domainDetails = taskDomainReservoir
	}

	return domainDetails, nil
}
