package mysql

import (
	"database/sql"

	"github.com/PalmaPedro/tania-core/src/devices/query"
	uuid "github.com/satori/go.uuid"
)

// CropQueryMysql :
type CropQueryMysql struct {
	DB *sql.DB
}

// NewCropQueryMysql :
func NewCropQueryMysql(db *sql.DB) query.CropQuery {
	return CropQueryMysql{DB: db}
}

// FindCropByID :
func (s CropQueryMysql) FindCropByID(uid uuid.UUID) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		rowsData := struct {
			UID     []byte
			BatchID string
		}{}
		crop := query.DeviceCropQueryResult{}

		err := s.DB.QueryRow(`SELECT UID, BATCH_ID
			FROM CROP_READ WHERE UID = ?`, uid.Bytes()).Scan(&rowsData.UID, &rowsData.BatchID)

		cropUID, err := uuid.FromBytes(rowsData.UID)
		if err != nil {
			result <- query.Result{Error: err}
		}

		crop.UID = cropUID
		crop.BatchID = rowsData.BatchID

		result <- query.Result{Result: crop}

		close(result)
	}()

	return result
}

