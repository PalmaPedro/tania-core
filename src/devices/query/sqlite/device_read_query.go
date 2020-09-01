package sqlite

import (
	"database/sql"
	//"strconv"
	"time"

	"github.com/PalmaPedro/tania-core/src/devices/domain"
	"github.com/PalmaPedro/tania-core/src/devices/query"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
	uuid "github.com/satori/go.uuid"
  "github.com/Tanibox/tania-core/src/helper/paginationhelper"
)

// DeviceReadQuerySqlite is used ...
type DeviceReadQuerySqlite struct {
	DB *sql.DB
}

// NewDeviceReadQuerySqlite is used ...
func NewDeviceReadQuerySqlite(s *sql.DB) query.DeviceReadQuery {
	return &DeviceReadQuerySqlite{DB: s}
}

// deviceReadQueryResult is used ...
type deviceReadQueryResult struct {
	UID                  string
	Title                string
	Description          string
	CreatedDate          string
	//DueDate              sql.NullString
	//CompletedDate        sql.NullString
	//CancelledDate        sql.NullString
	Priority             string
	Status               string
	DomainCode           string
	DomainDataMaterialID sql.NullString
	DomainDataAreaID     sql.NullString
	Category             string
	//IsDue                bool
	AssetID              sql.NullString
}

// FindAll is used ...
func (r DeviceReadQuerySqlite) FindAll(page, limit int) <-chan query.Result {
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

// FindByID is to find by ID
func (r DeviceReadQuerySqlite) FindByID(uid uuid.UUID) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		device := storage.DeviceRead{}

		rows, err := r.DB.Query(`SELECT * FROM DEVICE_READ WHERE UID = ?`, uid)
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

// FindTasksWithFilter is used ...
func (r DeviceReadQuerySqlite) FindTasksWithFilter(params map[string]string, page, limit int) <-chan query.Result {
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
			args = append(args, assetID)
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


// CountAll is used ...
func (r DeviceReadQuerySqlite) CountAll() <-chan query.Result {
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

// CountDevicesWithFilter is used ...
func (r DeviceReadQuerySqlite) CountDevicesWithFilter(params map[string]string) <-chan query.Result {
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
      args = append(args, assetID)
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

// populateQueryResult is used ....
func (r DeviceReadQuerySqlite) populateQueryResult(rows *sql.Rows) (storage.DeviceRead, error) {
	rowsData := deviceReadQueryResult{}

	err := rows.Scan(
		&rowsData.UID, &rowsData.Title, &rowsData.Description, &rowsData.CreatedDate,
		&rowsData.Priority, &rowsData.Status, &rowsData.DomainCode, &rowsData.DomainDataMaterialID,
		&rowsData.DomainDataAreaID,
		&rowsData.Category, &rowsData.AssetID,
	)

	if err != nil {
		return storage.DeviceRead{}, err
	}

	deviceUID, err := uuid.FromString(rowsData.UID)
	if err != nil {
		return storage.DeviceRead{}, err
	}

	createdDate, err := time.Parse(time.RFC3339, rowsData.CreatedDate)
	if err != nil {
		return storage.DeviceRead{}, err
	}

	/*
	var dueDate *time.Time
	if rowsData.DueDate.Valid && rowsData.DueDate.String != "" {
		d, err := time.Parse(time.RFC3339, rowsData.DueDate.String)
		if err != nil {
			return storage.TaskRead{}, err
		}

		dueDate = &d
	}

	var completedDate *time.Time
	if rowsData.CompletedDate.Valid && rowsData.CompletedDate.String != "" {
		d, err := time.Parse(time.RFC3339, rowsData.CompletedDate.String)
		if err != nil {
			return storage.TaskRead{}, err
		}

		completedDate = &d
	}

	var cancelledDate *time.Time
	if rowsData.CancelledDate.Valid && rowsData.CancelledDate.String != "" {
		d, err := time.Parse(time.RFC3339, rowsData.CancelledDate.String)
		if err != nil {
			return storage.TaskRead{}, err
		}

		cancelledDate = &d
	}*/

	var domainDetails domain.DeviceDomain
	switch rowsData.DomainCode {
	case domain.DeviceDomainAreaCode:
		materialID := (*uuid.UUID)(nil)
		if rowsData.DomainDataMaterialID.Valid && rowsData.DomainDataMaterialID.String != "" {
			uid, err := uuid.FromString(rowsData.DomainDataMaterialID.String)
			if err != nil {
				return storage.DeviceRead{}, err
			}
			materialID = &uid
		}
		domainDetails = domain.DeviceDomainArea{
			MaterialID: materialID,
		}

	case domain.DeviceDomainCropCode:
		materialID := (*uuid.UUID)(nil)
		areaID := (*uuid.UUID)(nil)
		if rowsData.DomainDataMaterialID.Valid && rowsData.DomainDataMaterialID.String != "" {
			uid, err := uuid.FromString(rowsData.DomainDataMaterialID.String)
			if err != nil {
				return storage.DeviceRead{}, err
			}
			materialID = &uid
		}
		if rowsData.DomainDataAreaID.Valid && rowsData.DomainDataAreaID.String != "" {
			uid, err := uuid.FromString(rowsData.DomainDataAreaID.String)
			if err != nil {
				return storage.DeviceRead{}, err
			}
			areaID = &uid
		}

		domainDetails = domain.DeviceDomainCrop{
			MaterialID: materialID,
			AreaID:     areaID}
	case domain.DeviceDomainFinanceCode:
		domainDetails = domain.DeviceDomainFinance{}
	case domain.DeviceDomainGeneralCode:
		domainDetails = domain.DeviceDomainGeneral{}
	case domain.DeviceDomainInventoryCode:
		domainDetails = domain.DeviceDomainInventory{}
	case domain.DeviceDomainReservoirCode:

		materialID := (*uuid.UUID)(nil)
		if rowsData.DomainDataMaterialID.Valid && rowsData.DomainDataMaterialID.String != "" {
			uid, err := uuid.FromString(rowsData.DomainDataMaterialID.String)
			if err != nil {
				return storage.DeviceRead{}, err
			}
			materialID = &uid
		}
		domainDetails = domain.DeviceDomainReservoir{
			MaterialID: materialID,
		}
	}

	var assetUID *uuid.UUID
	if rowsData.AssetID.Valid && rowsData.AssetID.String != "" {
		uid, err := uuid.FromString(rowsData.AssetID.String)
		if err != nil {
			return storage.DeviceRead{}, err
		}

		assetUID = &uid
	}

	return storage.DeviceRead{
		UID:           deviceUID,
		Title:         rowsData.Title,
		Description:   rowsData.Description,
		CreatedDate:   createdDate,
		//DueDate:       dueDate,
		//CompletedDate: completedDate,
		//CancelledDate: cancelledDate,
		//Priority:      rowsData.Priority,
		Status:        rowsData.Status,
		Domain:        rowsData.DomainCode,
		DomainDetails: domainDetails,
		Category:      rowsData.Category,
		//IsDue:         rowsData.IsDue,
		AssetID:       assetUID,
	}, nil
}


