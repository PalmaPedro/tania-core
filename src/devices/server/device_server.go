package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Tanibox/tania-core/config"
	assetsstorage "github.com/Tanibox/tania-core/src/assets/storage"
	"github.com/Tanibox/tania-core/src/eventbus"
	cropstorage "github.com/Tanibox/tania-core/src/growth/storage"
	"github.com/Tanibox/tania-core/src/helper/paginationhelper"
	"github.com/Tanibox/tania-core/src/helper/structhelper"
	"github.com/Tanibox/tania-core/src/tasks/domain"
	service "github.com/Tanibox/tania-core/src/tasks/domain/service"
	"github.com/Tanibox/tania-core/src/tasks/query"
	queryInMem "github.com/Tanibox/tania-core/src/tasks/query/inmemory"
	queryMysql "github.com/Tanibox/tania-core/src/tasks/query/mysql"
	querySqlite "github.com/Tanibox/tania-core/src/tasks/query/sqlite"
	"github.com/Tanibox/tania-core/src/tasks/repository"
	repoInMem "github.com/Tanibox/tania-core/src/tasks/repository/inmemory"
	repoMysql "github.com/Tanibox/tania-core/src/tasks/repository/mysql"
	repoSqlite "github.com/Tanibox/tania-core/src/tasks/repository/sqlite"
	"github.com/Tanibox/tania-core/src/tasks/storage"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"

)