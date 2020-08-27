package repository


import (
	"database/sql"

	"github.com/PalmaPedro/tania-core/src/devices/domain"
	"github.com/PalmaPedro/tania-core/src/devices/repository"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
)

// DeviceReadRepositoryMysql is used ...
type DeviceReadRepositoryMysql struct {
	DB *sql.DB
}

// NewDeviceReadRepositoryMysql is used ...
func NewDeviceReadRepositoryMysql(s *sql.DB) repository.DeviceReadRepository {
	return &DeviceReadRepositoryMysql{DB: s}
}

// Save is used ...
func (f *DeviceReadRepositoryMysql) Save(deviceRead *storage.DeviceRead) <-chan error {
	result := make(chan error)

	go func() {
		var domainDataMaterialID []byte
		var domainDataAreaID []byte
		switch v := deviceRead.DomainDetails.(type) {
		case domain.DeviceDomainCrop:
			if v.MaterialID != nil {
				domainDataMaterialID = v.MaterialID.Bytes()
			}
			if v.AreaID != nil {
				domainDataAreaID = v.AreaID.Bytes()
			}
		}

		var assetID []byte
		if deviceRead.AssetID != nil {
			assetID = deviceRead.AssetID.Bytes()
		}

		res, err := f.DB.Exec(`UPDATE DEVICE_READ SET
			TITLE = ?, DESCRIPTION = ?, STATUS = ?,
			DOMAIN_CODE = ?, DOMAIN_DATA_MATERIAL_ID = ?, DOMAIN_DATA_AREA_ID = ?,
			CATEGORY = ?, ASSET_ID = ?
			WHERE UID = ?`,
			deviceRead.Title, deviceRead.Description, deviceRead.CreatedDate,
			deviceRead.Status,
			deviceRead.Domain,domainDataAreaID,
			deviceRead.Category, assetID,
			deviceRead.UID.Bytes())

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
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				deviceRead.UID.Bytes(), deviceRead.Title, deviceRead.Description, deviceRead.CreatedDate,
				deviceRead.Status,
				deviceRead.Domain, domainDataMaterialID, domainDataAreaID,
				deviceRead.Category, assetID)

			if err != nil {
				result <- err
			}
		}

		result <- nil
		close(result)
	}()

	return result
}