package mysql


import (
	"database/sql"
	//"strconv"
	"time"

	"github.com/Tanibox/tania-core/src/helper/paginationhelper"
	"github.com/PalmaPedro/tania-core/src/devices/domain"
	"github.com/PalmaPedro/tania-core/src/devices/query"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
	uuid "github.com/satori/go.uuid"
)

// DeviceReadQueryMysql :
type DeviceReadQueryMysql struct {
	DB *sql.DB
}

// NewDeviceReadQueryMysql :
func NewDeviceReadQueryMysql(s *sql.DB) query.DeviceReadQuery {
	return &DeviceReadQueryMysql{DB: s}
}

type deviceReadQueryResult struct {
	UID                  []byte
	Title                string
	Description          string
	CreatedDate          time.Time
	Status               string
	DomainCode           string
	DomainDataMaterialID uuid.NullUUID
	DomainDataAreaID     uuid.NullUUID
	Category             string
	AssetID              uuid.NullUUID
}

// FindAll :
func (r DeviceReadQueryMysql) FindAll(page, limit int) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		devices := []storage.DeviceRead{}

		sql := `SELECT * FROM DEVICE_READ ORDER BY CREATED_DATE DESC`
		var args []interface{}

		if page != 0 && limit != 0 {
			sql += " LIMIT ? OFFSET ?"
			offset := paginationhelper.CalculatePageToOffset(page, limit)
			args = append(args, limit, offset)
		}

		rows, err := r.DB.Query(sql, args...)
		if err != nil {
			result <- query.Result{Error: err}
		}

		for rows.Next() {
			deviceRead, err := r.populateQueryResult(rows)
			if err != nil {
				result <- query.Result{Error: err}
			}

			devices = append(devices, deviceRead)
		}

		result <- query.Result{Result: devices}

		close(result)
	}()

	return result
}

// FindByID :
func (r DeviceReadQueryMysql) FindByID(uid uuid.UUID) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		device := storage.DeviceRead{}

		rows, err := r.DB.Query(`SELECT * FROM DEVICE_READ WHERE UID = ?`, uid.Bytes())
		if err != nil {
			result <- query.Result{Error: err}
		}

		for rows.Next() {
			deviceRead, err := r.populateQueryResult(rows)
			if err != nil {
				result <- query.Result{Error: err}
			}

			device = deviceRead
		}

		result <- query.Result{Result: device}
		close(result)
	}()

	return result
}

// FindDevicesWithFilter :
func (r DeviceReadQueryMysql) FindDevicesWithFilter(params map[string]string, page, limit int) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		devices := []storage.DeviceRead{}

		sql := "SELECT * FROM DEVICE_READ WHERE 1 = 1"
		var args []interface{}
		/*
		if value, _ := params["is_due"]; value != "" {
			b, _ := strconv.ParseBool(value)
			sql += " AND IS_DUE = ? "
			args = append(args, b)
		}
		start, _ := params["due_start"]
		end, _ := params["due_end"]
		if start != "" && end != "" {
			sql += " AND DUE_DATE BETWEEN ? AND ? "
			args = append(args, start)
			args = append(args, end)
		}
		if value, _ := params["priority"]; value != "" {
			sql += " AND PRIORITY = ? "
			args = append(args, value)
		}*/
		if value, _ := params["status"]; value != "" {
			sql += " AND STATUS = ? "
			args = append(args, value)
		}
		if value, _ := params["domain"]; value != "" {
			sql += " AND DOMAIN_CODE = ? "
			args = append(args, value)
		}
		if value, _ := params["category"]; value != "" {
			sql += " AND CATEGORY = ? "
			args = append(args, value)
		}
		if value, _ := params["asset_id"]; value != "" {
			assetID, _ := uuid.FromString(value)
			sql += " AND ASSET_ID = ? "
			args = append(args, assetID.Bytes())
		}

		if page != 0 && limit != 0 {
			sql += " LIMIT ? OFFSET ?"
			offset := paginationhelper.CalculatePageToOffset(page, limit)
			args = append(args, limit, offset)
		}

		rows, err := r.DB.Query(sql, args...)
		if err != nil {
			result <- query.Result{Error: err}
		}

		for rows.Next() {
			deviceRead, err := r.populateQueryResult(rows)
			if err != nil {
				result <- query.Result{Error: err}
			}

			devices = append(devices, deviceRead)
		}

		result <- query.Result{Result: devices}

		close(result)
	}()

	return result
}

// CountAll :
func (r DeviceReadQueryMysql) CountAll() <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		total := 0
		var params []interface{}

		sql := "SELECT COUNT(UID) FROM DEVICE_READ"

		err := r.DB.QueryRow(sql, params...).Scan(&total)
		if err != nil {
			result <- query.Result{Error: err}
		}

		result <- query.Result{Result: total}
		close(result)
	}()

	return result
}

// CountDevicesWithFilter :
func (r DeviceReadQueryMysql) CountDevicesWithFilter(params map[string]string) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		total := 0

		sql := "SELECT COUNT(UID) FROM DEVICE_READ WHERE 1 = 1"
		var args []interface{}
		/*
		if value, _ := params["is_due"]; value != "" {
			b, _ := strconv.ParseBool(value)
			sql += " AND IS_DUE = ? "
			args = append(args, b)
		}
		start, _ := params["due_start"]
		end, _ := params["due_end"]
		if start != "" && end != "" {
			sql += " AND DUE_DATE BETWEEN ? AND ? "
			args = append(args, start)
			args = append(args, end)
		}
		if value, _ := params["priority"]; value != "" {
			sql += " AND PRIORITY = ? "
			args = append(args, value)
		}*/
		if value, _ := params["status"]; value != "" {
			sql += " AND STATUS = ? "
			args = append(args, value)
		}
		if value, _ := params["domain"]; value != "" {
			sql += " AND DOMAIN_CODE = ? "
			args = append(args, value)
		}
		if value, _ := params["category"]; value != "" {
			sql += " AND CATEGORY = ? "
			args = append(args, value)
		}
		if value, _ := params["asset_id"]; value != "" {
			assetID, _ := uuid.FromString(value)
			sql += " AND ASSET_ID = ? "
			args = append(args, assetID.Bytes())
		}

		err := r.DB.QueryRow(sql, args...).Scan(&total)
		if err != nil {
			result <- query.Result{Error: err}
		}

		result <- query.Result{Result: total}
		close(result)
	}()

	return result
}

// populateQueryResult :
func (r DeviceReadQueryMysql) populateQueryResult(rows *sql.Rows) (storage.DeviceRead, error) {
	rowsData := deviceReadQueryResult{}

	err := rows.Scan(
		&rowsData.UID, &rowsData.Title, &rowsData.Description, &rowsData.CreatedDate,
		&rowsData.DomainDataAreaID, &rowsData.Category, &rowsData.AssetID,
	)

	if err != nil {
		return storage.DeviceRead{}, err
	}

	deviceUID, err := uuid.FromBytes(rowsData.UID)
	if err != nil {
		return storage.DeviceRead{}, err
	}

	var domainDetails domain.DeviceDomain
	switch rowsData.DomainCode {
	case domain.DeviceDomainAreaCode:
		domainDetails = domain.DeviceDomainArea{}
	case domain.DeviceDomainCropCode:
		var materialID *uuid.UUID
		var areaID *uuid.UUID

		if rowsData.DomainDataMaterialID.Valid {
			materialID = &rowsData.DomainDataMaterialID.UUID
		}
		if rowsData.DomainDataAreaID.Valid {
			areaID = &rowsData.DomainDataAreaID.UUID
		}

		domainDetails = domain.DeviceDomainCrop{
			MaterialID: materialID,
			AreaID:     areaID,
		}

	//case domain.DeviceDomainFinanceCode:
	//	domainDetails = domain.DeviceDomainFinance{}
	case domain.DeviceDomainGeneralCode:
		domainDetails = domain.DeviceDomainGeneral{}
	}

	assetUID := &uuid.UUID{}
	if rowsData.AssetID.Valid {
		assetUID = &rowsData.AssetID.UUID
	}
	/*
	isDue := false
	if rowsData.IsDue == 1 {
		isDue = true
	}*/

	return storage.DeviceRead{
		UID:           deviceUID,
		Title:         rowsData.Title,
		Description:   rowsData.Description,
		CreatedDate:   rowsData.CreatedDate,
		Status:        rowsData.Status,
		Domain:        rowsData.DomainCode,
		DomainDetails: domainDetails,
		Category:      rowsData.Category,
		AssetID:       assetUID,
	}, nil
}
