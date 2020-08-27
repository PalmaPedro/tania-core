package mysql

import (
	"database/sql"

	"github.com/PalmaPedro/tania-core/src/devices/query"
	uuid "github.com/satori/go.uuid"
)

// AreaQueryMysql : 
type AreaQueryMysql struct {
	DB *sql.DB
}

// NewAreaQueryMysql :
func NewAreaQueryMysql(db *sql.DB) query.AreaQuery {
	return AreaQueryMysql{DB: db}
}

// FindByID : 
func (s AreaQueryMysql) FindByID(uid uuid.UUID) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		rowsData := struct {
			UID  []byte
			Name string
		}{}
		area := query.DeviceAreaQueryResult{}

		err := s.DB.QueryRow(`SELECT UID, NAME
			FROM AREA_READ WHERE UID = ?`, uid.Bytes()).Scan(&rowsData.UID, &rowsData.Name)

		areaUID, err := uuid.FromBytes(rowsData.UID)
		if err != nil {
			result <- query.Result{Error: err}
		}

		area.UID = areaUID
		area.Name = rowsData.Name

		result <- query.Result{Result: area}

		close(result)
	}()

	return result
}

