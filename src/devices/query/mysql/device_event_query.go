package mysql

import (
	"database/sql"
	"encoding/json"
	"time"


	"github.com/PalmaPedro/tania-core/src/devices/decoder"
	"github.com/PalmaPedro/tania-core/src/devices/query"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
	uuid "github.com/satori/go.uuid"
)

// DeviceEventQueryMysql :
type DeviceEventQueryMysql struct {
	DB *sql.DB
}

// NewDeviceEventQueryMysql :
func NewDeviceEventQueryMysql(db *sql.DB) query.DeviceEventQuery {
	return &DeviceEventQueryMysql{DB: db}
}

// FindAllByDeviceID :
func (f *DeviceEventQueryMysql) FindAllByDeviceID(uid uuid.UUID) <-chan query.QueryResult {
	result := make(chan query.QueryResult)

	go func() {
		events := []storage.DeviceEvent{}

		rows, err := f.DB.Query("SELECT * FROM DEVICE_EVENT WHERE DEVICE_UID = ? ORDER BY VERSION ASC", uid.Bytes())
		if err != nil {
			result <- query.QueryResult{Error: err}
		}

		rowsData := struct {
			ID          int
			TaskUID     []byte
			Version     int
			CreatedDate time.Time
			Event       []byte
		}{}

		for rows.Next() {
			rows.Scan(&rowsData.ID, &rowsData.DeviceUID, &rowsData.Version, &rowsData.CreatedDate, &rowsData.Event)

			wrapper := decoder.DeviceEventWrapper{}
			json.Unmarshal(rowsData.Event, &wrapper)

			deviceUID, err := uuid.FromBytes(rowsData.DeviceUID)
			if err != nil {
				result <- query.QueryResult{Error: err}
			}

			events = append(events, storage.DeviceEvent{
				DeviceUID:     deviceUID,
				Version:     rowsData.Version,
				CreatedDate: rowsData.CreatedDate,
				Event:       wrapper.Data,
			})
		}

		result <- query.QueryResult{Result: events}
		close(result)
	}()

	return result
}



