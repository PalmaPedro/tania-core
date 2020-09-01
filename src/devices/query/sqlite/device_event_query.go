package sqlite

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/PalmaPedro/tania-core/src/devices/decoder"

	"github.com/PalmaPedro/tania-core/src/devices/query"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
	uuid "github.com/satori/go.uuid"
)

// DeviceEventQuerySqlite is used ...
type DeviceEventQuerySqlite struct {
	DB *sql.DB
}

// NewDeviceEventQuerySqlite is used ...
func NewDeviceEventQuerySqlite(db *sql.DB) query.DeviceEventQuery {
	return &DeviceEventQuerySqlite{DB: db}
}

// FindAllByDeviceID is used ...
func (f *DeviceEventQuerySqlite) FindAllByDeviceID(uid uuid.UUID) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		events := []storage.DeviceEvent{}

		rows, err := f.DB.Query("SELECT * FROM DEVICE_EVENT WHERE DEVICE_UID = ? ORDER BY VERSION ASC", uid)
		if err != nil {
			result <- query.Result{Error: err}
		}

		rowsData := struct {
			ID          int
			DeviceUID     string
			Version     int
			CreatedDate string
			Event       []byte
		}{}

		for rows.Next() {
			rows.Scan(&rowsData.ID, &rowsData.DeviceUID, &rowsData.Version, &rowsData.CreatedDate, &rowsData.Event)

			wrapper := decoder.DeviceEventWrapper{}
			json.Unmarshal(rowsData.Event, &wrapper)

			deviceUID, err := uuid.FromString(rowsData.DeviceUID)
			if err != nil {
				result <- query.Result{Error: err}
			}

			createdDate, err := time.Parse(time.RFC3339, rowsData.CreatedDate)
			if err != nil {
				result <- query.Result{Error: err}
			}

			events = append(events, storage.DeviceEvent{
				DeviceUID:     deviceUID,
				Version:     rowsData.Version,
				CreatedDate: createdDate,
				Event:       wrapper.Data,
			})
		}

		result <- query.Result{Result: events}
		close(result)
	}()

	return result
}