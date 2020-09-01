package repository

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/Tanibox/tania-core/src/helper/structhelper"
	"github.com/PalmaPedro/tania-core/src/devices/decoder"
	"github.com/PalmaPedro/tania-core/src/devices/repository"
	uuid "github.com/satori/go.uuid"
)

// DeviceEventRepositorySqlite is used ...
type DeviceEventRepositorySqlite struct {
	DB *sql.DB
}

// NewDeviceEventRepositorySqlite is used
func NewDeviceEventRepositorySqlite(s *sql.DB) repository.DeviceEventRepository {
	return &DeviceEventRepositorySqlite{DB: s}
}

// Save is used to save a new device into the database, device_event table 
func (s *DeviceEventRepositorySqlite) Save(uid uuid.UUID, latestVersion int, events []interface{}) <-chan error {
	result := make(chan error)

	go func() {
		for _, v := range events {
			latestVersion++

			stmt, err := s.DB.Prepare(`INSERT INTO DEVICE_EVENT
				(DEVICE_UID, VERSION, CREATED_DATE, EVENT)
				VALUES (?, ?, ?, ?)`)

			if err != nil {
				result <- err
			}

			e, err := json.Marshal(decoder.InterfaceWrapper{
				Name: structhelper.GetName(v),
				Data: v,
			})

			if err != nil {
				panic(err)
			}

			_, err = stmt.Exec(uid, latestVersion, time.Now().Format(time.RFC3339), e)
			if err != nil {
				result <- err
			}
		}

		result <- nil
		close(result)
	}()

	return result
}