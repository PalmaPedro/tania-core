package mysql

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/PalmaPedro/tania-core/src/devices/query"
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

