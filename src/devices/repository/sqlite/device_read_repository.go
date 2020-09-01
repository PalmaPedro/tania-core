package repository

import (
	"database/sql"
	"github.com/PalmaPedro/tania-core/src/devices/domain"
	"github.com/PalmaPedro/tania-core/src/devices/repository"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
	uuid "github.com/satori/go.uuid"
	"time"
)

// DeviceReadRepositorySqlite is used ...
type DeviceReadRepositorySqlite struct {
	DB *sql.DB
}

// NewDeviceReadRepositorySqlite is used ..
func NewDeviceReadRepositorySqlite(s *sql.DB) repository.DeviceReadRepository {
	return &DeviceReadRepositorySqlite{DB: s}
}

// Save is used ...
func (f *DeviceReadRepositorySqlite) Save(deviceRead *storage.DeviceRead) <-chan error {
	result := make(chan error)

	go func() {
		/*
		var dueDate *string
		if taskRead.DueDate != nil && !taskRead.DueDate.IsZero() {
			d := taskRead.DueDate.Format(time.RFC3339)
			dueDate = &d
		}

		var completedDate *string
		if taskRead.CompletedDate != nil && !taskRead.CompletedDate.IsZero() {
			d := taskRead.CompletedDate.Format(time.RFC3339)
			completedDate = &d
		}

		var cancelledDate *string
		if taskRead.CancelledDate != nil && !taskRead.CancelledDate.IsZero() {
			d := taskRead.CancelledDate.Format(time.RFC3339)
			cancelledDate = &d
		}
		*/
		var domainDataMaterialID *uuid.UUID
		var domainDataAreaID *uuid.UUID
		switch v := deviceRead.DomainDetails.(type) {
		case domain.DeviceDomainArea:
			domainDataMaterialID = v.MaterialID
		case domain.DeviceDomainCrop:
			domainDataMaterialID = v.MaterialID
			domainDataAreaID = v.AreaID
		case domain.DeviceDomainReservoir:
			domainDataMaterialID = v.MaterialID
		}

		res, err := f.DB.Exec(`UPDATE DEVICE_READ SET
			TITLE = ?, DESCRIPTION = ?, CREATED_DATE = ?, STATUS = ?,
			DOMAIN_CODE = ?, DOMAIN_DATA_MATERIAL_ID = ?, DOMAIN_DATA_AREA_ID = ?,
			CATEGORY = ?, ASSET_ID = ?
			WHERE UID = ?`,
			deviceRead.Title, deviceRead.Description, deviceRead.CreatedDate.Format(time.RFC3339), deviceRead.Status,
			deviceRead.Domain, domainDataMaterialID, domainDataAreaID, deviceRead.Category, deviceRead.AssetID,
			deviceRead.UID)

		if err != nil {
			result <- err
		}

		rowsAffected := int64(0)
		if res != nil {
			rowsAffected, err = res.RowsAffected()
			if err != nil {
				result <- err
			}
		}

		if rowsAffected == 0 {
			_, err := f.DB.Exec(`INSERT INTO DEVICE_READ (
				UID, TITLE, DESCRIPTION, CREATED_DATE, STATUS,
				DOMAIN_CODE, DOMAIN_DATA_MATERIAL_ID, DOMAIN_DATA_AREA_ID, CATEGORY, ASSET_ID)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				deviceRead.UID, deviceRead.Title, deviceRead.Description, deviceRead.CreatedDate.Format(time.RFC3339),
				deviceRead.Status, deviceRead.Domain, domainDataMaterialID, domainDataAreaID,
				deviceRead.Category, deviceRead.AssetID)

			if err != nil {
				result <- err
			}
		}

		result <- nil
		close(result)
	}()

	return result
}
