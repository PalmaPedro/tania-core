package server

import (
	"database/sql"
	"net/http"
	//"time"

	"github.com/Tanibox/tania-core/config"
	"github.com/Tanibox/tania-core/src/eventbus"
	"github.com/PalmaPedro/tania-core/src/devices/query"
	"github.com/PalmaPedro/tania-core/src/devices/repository"
	"github.com/PalmaPedro/tania-core/src/devices/domain"
	"github.com/Tanibox/tania-core/src/helper/structhelper"
	service "github.com/PalmaPedro/tania-core/src/devices/domain/service"
	assetsstorage "github.com/Tanibox/tania-core/src/assets/storage"
	cropstorage "github.com/Tanibox/tania-core/src/growth/storage"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
	repoInMem "github.com/PalmaPedro/tania-core/src/devices/repository/inmemory"
	queryInMem "github.com/PalmaPedro/tania-core/src/devices/query/inmemory"
	queryMysql "github.com/PalmaPedro/tania-core/src/devices/query/mysql"
	repoMysql "github.com/PalmaPedro/tania-core/src/devices/repository/mysql"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// DeviceServer ties the routes and handlers with injected dependencies
type DeviceServer struct {
	DeviceEventRepo  repository.DeviceEventRepository
	DeviceReadRepo   repository.DeviceReadRepository
	DeviceEventQuery query.DeviceEventQuery
	DeviceReadQuery  query.DeviceReadQuery
	DeviceService    domain.DeviceService
	EventBus       	 eventbus.TaniaEventBus
}

// NewDeviceServer initializes DeviceServer's dependencies and create new DeviceServer struct
func NewDeviceServer(
	db *sql.DB,
	bus eventbus.TaniaEventBus,
	cropStorage *cropstorage.CropReadStorage,
	areaStorage *assetsstorage.AreaReadStorage,
	materialStorage *assetsstorage.MaterialReadStorage,
	deviceEventStorage *storage.DeviceEventStorage,
	deviceReadStorage *storage.DeviceReadStorage) (*DeviceServer, error) {

	deviceServer := &DeviceServer{
		EventBus: bus,
	}

	switch *config.Config.TaniaPersistenceEngine {
	case config.DB_INMEMORY:
		deviceServer.DeviceEventRepo = repoInMem.NewDeviceEventRepositoryInMemory(deviceEventStorage)
		deviceServer.DeviceReadRepo = repoInMem.NewDeviceReadRepositoryInMemory(deviceReadStorage)

		deviceServer.DeviceEventQuery = queryInMem.NewDeviceEventQueryInMemory(deviceEventStorage)
		deviceServer.DeviceReadQuery = queryInMem.NewDeviceReadQueryInMemory(deviceReadStorage)

		cropQuery := queryInMem.NewCropQueryInMemory(cropStorage)
		areaQuery := queryInMem.NewAreaQueryInMemory(areaStorage)
		materialReadQuery := queryInMem.NewMaterialQueryInMemory(materialStorage)

		deviceServer.DeviceService = service.DeviceServiceSqLLite{
			CropQuery:      cropQuery,
			AreaQuery:      areaQuery,
			MaterialQuery:  materialReadQuery,
		}

	case config.DB_MYSQL:
		deviceServer.DeviceEventRepo = repoMysql.NewDeviceEventRepositoryMysql(db)
		deviceServer.DeviceReadRepo = repoMysql.NewDeviceReadRepositoryMysql(db)

		deviceServer.DeviceEventQuery = queryMysql.NewDeviceEventQueryMysql(db)
		deviceServer.DeviceReadQuery = queryMysql.NewDeviceReadQueryMysql(db)

		cropQuery := queryMysql.NewCropQueryMysql(db)
		areaQuery := queryMysql.NewAreaQueryMysql(db)
		materialReadQuery := queryMysql.NewMaterialQueryMysql(db)

		deviceServer.DeviceService = service.DeviceServiceSqLLite{
			CropQuery:      cropQuery,
			AreaQuery:      areaQuery,
			MaterialQuery:  materialReadQuery,
		}

	}

	deviceServer.InitSubscriber()

	return deviceServer, nil
}

// InitSubscriber defines the mapping of which event this domain listen with their handler
func (s *DeviceServer) InitSubscriber() {
	
}

// Mount defines the DeviceServer's endpoints with its handlers
func (s *DeviceServer) Mount(g *echo.Group) {
	g.POST("", s.SaveDevice)
	//g.GET("", s.FindAllDevices)
	//g.GET("/search", s.FindFilteredDevices)
	//g.GET("/:id", s.FindDeviceByID)
	//g.PUT("/:id", s.UpdateDevice)
	//g.PUT("/:id/cancel", s.CancelDevice)
}

// SaveDevice is a DeviceServer's handler to save new Device
func (s *DeviceServer) SaveDevice(c echo.Context) error {

	data := make(map[string]storage.DeviceRead)

	/*formDate := c.FormValue("due_date")
	duePtr := (*time.Time)(nil)
	if len(formDate) != 0 {
		dueDate, err := time.Parse(time.RFC3339Nano, formDate)

		if err != nil {
			return Error(c, err)
		}
		duePtr = &dueDate
	}*/

	assetID := c.FormValue("asset_id")
	assetIDPtr := (*uuid.UUID)(nil)
	if len(assetID) != 0 {
		assetID, err := uuid.FromString(assetID)
		if err != nil {
			return Error(c, err)
		}
		assetIDPtr = &assetID
	}

	domaincode := c.FormValue("domain")

	domaindevice, err := s.CreateDeviceDomainByCode(domaincode, c)

	if err != nil {
		return Error(c, err)
	}

	device, err := domain.CreateDevice(
		s.DeviceService,
		c.FormValue("title"),
		c.FormValue("description"),
		//duePtr,
		//c.FormValue("priority"),
		domaindevice,
		c.FormValue("category"),
		assetIDPtr)

	if err != nil {
		return Error(c, err)
	}

	err = <-s.DeviceEventRepo.Save(device.UID, 0, device.UncommittedChanges)
	if err != nil {
		return Error(c, err)
	}

	// Trigger Events
	s.publishUncommittedEvents(device)

	deviceRead := MapDeviceToDeviceRead(device)
	s.AppendDeviceDomainDetails(deviceRead)

	data["data"] = *deviceRead

	return c.JSON(http.StatusOK, data)
}

// CreateDeviceDomainByCode is used...
func (s *DeviceServer) CreateDeviceDomainByCode(domaincode string, c echo.Context) (domain.DeviceDomain, error) {
	domainvalue := domaincode
	if domainvalue == "" {
		return nil, NewRequestValidationError(required, "domain")
	}

	switch domainvalue {
	case domain.DeviceDomainAreaCode:

		category := c.FormValue("category")
		materialID := c.FormValue("material_id")

		materialPtr := (*uuid.UUID)(nil)
		if len(materialID) != 0 {
			uid, err := uuid.FromString(materialID)
			if err != nil {
				return domain.DeviceDomainArea{}, err
			}
			materialPtr = &uid
		}

		return domain.CreateDeviceDomainArea(s.DeviceService, category, materialPtr)
	case domain.DeviceDomainCropCode:

		category := c.FormValue("category")
		materialID := c.FormValue("material_id")
		areaID := c.FormValue("area_id")

		materialPtr := (*uuid.UUID)(nil)
		if len(materialID) != 0 {
			uid, err := uuid.FromString(materialID)
			if err != nil {
				return domain.DeviceDomainCrop{}, err
			}
			materialPtr = &uid
		}

		areaPtr := (*uuid.UUID)(nil)
		if len(areaID) != 0 {
			uid, err := uuid.FromString(areaID)
			if err != nil {
				return domain.DeviceDomainCrop{}, err
			}
			areaPtr = &uid
		}

		return domain.CreateDeviceDomainCrop(s.DeviceService, category, materialPtr, areaPtr)
	case domain.DeviceDomainFinanceCode:
		return domain.CreateDeviceDomainFinance()
	case domain.DeviceDomainGeneralCode:
		return domain.CreateDeviceDomainGeneral()
	case domain.DeviceDomainInventoryCode:
		return domain.CreateDeviceDomainInventory()
	case domain.DeviceDomainReservoirCode:

		category := c.FormValue("category")
		materialID := c.FormValue("material_id")

		materialPtr := (*uuid.UUID)(nil)
		if len(materialID) != 0 {
			uid, err := uuid.FromString(materialID)
			if err != nil {
				return domain.DeviceDomainReservoir{}, err
			}
			materialPtr = &uid
		}
		return domain.CreateDeviceDomainReservoir(s.DeviceService, category, materialPtr)
	default:
		return nil, NewRequestValidationError(invalidOption, "domain")
	}
}

func (s *DeviceServer) publishUncommittedEvents(entity interface{}) error {

	switch e := entity.(type) {
	case *domain.Device:
		for _, v := range e.UncommittedChanges {
			name := structhelper.GetName(v)
			s.EventBus.Publish(name, v)
		}
	default:
	}

	return nil
}

// AppendDeviceDomainDetails is used ...
func (s *DeviceServer) AppendDeviceDomainDetails(device *storage.DeviceRead) error {

	switch device.Domain {
	case domain.DeviceDomainAreaCode:
		materialID := device.DomainDetails.(domain.DeviceDomainArea).MaterialID
		if materialID != nil {
			materialResult := s.DeviceService.FindMaterialByID(*materialID)
			materialQueryResult, ok := materialResult.Result.(query.DeviceMaterialQueryResult)

			if !ok {
				return echo.NewHTTPError(http.StatusBadRequest, "Internal server error")
			}
			device.DomainDetails = &storage.DeviceDomainDetailedArea{
				MaterialID:           &materialQueryResult.UID,
				MaterialName:         materialQueryResult.Name,
				MaterialType:         materialQueryResult.TypeCode,
				MaterialDetailedType: materialQueryResult.DetailedTypeCode,
			}
		}

	case domain.DeviceDomainCropCode:
		material := (*storage.DeviceDomainCropMaterial)(nil)
		area := (*storage.DeviceDomainCropArea)(nil)
		crop := (*storage.DeviceDomainCropBatch)(nil)

		
		materialID := device.DomainDetails.(domain.DeviceDomainCrop).MaterialID
		if materialID != nil {
			materialResult := s.DeviceService.FindMaterialByID(*materialID)
			materialQueryResult, ok := materialResult.Result.(query.DeviceMaterialQueryResult)

			if !ok {
				return echo.NewHTTPError(http.StatusBadRequest, "Internal server error")
			}
			material = &storage.DeviceDomainCropMaterial{
				MaterialID:           &materialQueryResult.UID,
				MaterialName:         materialQueryResult.Name,
				MaterialType:         materialQueryResult.TypeCode,
				MaterialDetailedType: materialQueryResult.DetailedTypeCode,
			}
		}

		areaID := device.DomainDetails.(domain.DeviceDomainCrop).AreaID
		if areaID != nil {
			areaResult := s.DeviceService.FindAreaByID(*areaID)
			areaQueryResult, ok := areaResult.Result.(query.DeviceAreaQueryResult)

			if !ok {
				return echo.NewHTTPError(http.StatusBadRequest, "Internal server error")
			}
			area = &storage.DeviceDomainCropArea{
				AreaID:   &areaQueryResult.UID,
				AreaName: areaQueryResult.Name,
			}
		}
		
		cropID := device.AssetID
		if cropID != nil {
			cropResult := s.DeviceService.FindCropByID(*cropID)
			cropQueryResult, ok := cropResult.Result.(query.DeviceCropQueryResult)

			if !ok {
				return echo.NewHTTPError(http.StatusBadRequest, "Internal server error")
			}
			crop = &storage.DeviceDomainCropBatch{
				CropID:      &cropQueryResult.UID,
				CropBatchID: cropQueryResult.BatchID,
			}
		}
		device.DomainDetails = &storage.DeviceDomainDetailedCrop{
			Material: material,
			Area:     area,
			Crop:     crop,
		}
	
	
	case domain.DeviceDomainReservoirCode:

		materialID := device.DomainDetails.(domain.DeviceDomainReservoir).MaterialID
		if materialID != nil {
			materialResult := s.DeviceService.FindMaterialByID(*materialID)
			materialQueryResult, ok := materialResult.Result.(query.DeviceMaterialQueryResult)

			if !ok {
				return echo.NewHTTPError(http.StatusBadRequest, "Internal server error")
			}
			device.DomainDetails = &storage.DeviceDomainDetailedReservoir{
				MaterialID:           &materialQueryResult.UID,
				MaterialName:         materialQueryResult.Name,
				MaterialType:         materialQueryResult.TypeCode,
				MaterialDetailedType: materialQueryResult.DetailedTypeCode,
			}
		}
	}

	return nil
}
