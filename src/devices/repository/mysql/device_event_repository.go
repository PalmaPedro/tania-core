package repository

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/PalmaPedro/tania-core/src/helper/structhelper"
	"github.com/PalmaPedro/tania-core/src/devices/decoder"
	"github.com/PalmaPedro/tania-core/src/devices/repository"
	uuid "github.com/satori/go.uuid"
)

// DeviceEventRepositoryMysql is used ...
type DeviceEventRepositoryMysql struct {
	DB *sql.DB
}

// NewDeviceEventRepositoryMysql is used...
func NewDeviceEventRepositoryMysql(s *sql.DB) repository.DeviceEventRepository {
	return &DeviceEventRepositoryMysql{DB: s}
}

// Save is used ...
func (s *DeviceEventRepositoryMysql) Save(uid uuid.UUID, latestVersion int, events []interface{}) <-chan error {
	result := make(chan error)

	go func() {
		for _, v := range events {
			latestVersion++

			stmt, err := s.DB.Prepare(`INSERT INTO DEVICE_EVENT
				(TASK_UID, VERSION, CREATED_DATE, EVENT)
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

			_, err = stmt.Exec(uid.Bytes(), latestVersion, time.Now(), e)
			if err != nil {
				result <- err
			}
		}

		result <- nil
		close(result)
	}()

	return result
}