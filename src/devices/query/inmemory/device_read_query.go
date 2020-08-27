package inmemory

import (
	"github.com/PalmaPedro/tania-core/src/devices/query"
	"github.com/PalmaPedro/tania-core/src/devices/storage"
	uuid "github.com/satori/go.uuid"
	//"strconv"
	//"time"
)

// DeviceReadQueryInMemory is used ...
type DeviceReadQueryInMemory struct {
	Storage *storage.DeviceReadStorage
}

// NewDeviceReadQueryInMemory is used ...
func NewDeviceReadQueryInMemory(s *storage.DeviceReadStorage) query.DeviceReadQuery {
	return &DeviceReadQueryInMemory{Storage: s}
}

// FindAll is used ...
func (r DeviceReadQueryInMemory) FindAll(page, limit int) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		r.Storage.Lock.RLock()
		defer r.Storage.Lock.RUnlock()

		devices := []storage.DeviceRead{}

		for _, val := range r.Storage.DeviceReadMap {
			devices = append(devices, val)
		}

		if limit != 0 {
			devices = devices[:limit]
		}

		result <- query.Result{Result: devices}

		close(result)
	}()

	return result
}

// FindByID is to find by ID
func (r DeviceReadQueryInMemory) FindByID(uid uuid.UUID) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		r.Storage.Lock.RLock()
		defer r.Storage.Lock.RUnlock()

		result <- query.Result{Result: r.Storage.DeviceReadMap[uid]}

		close(result)
	}()

	return result
}

// FindDevicesWithFilter is used ...
func (r DeviceReadQueryInMemory) FindDevicesWithFilter(params map[string]string, page, limit int) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		r.Storage.Lock.RLock()
		defer r.Storage.Lock.RUnlock()

		devices := []storage.DeviceRead{}
		for _, val := range r.Storage.DeviceReadMap {
			isMatch := true
/*
			// Is Due
			if value, _ := params["is_due"]; value != "" {
				b, _ := strconv.ParseBool(value)
				if val.IsDue != b {
					is_match = false
				}
			}
*/

			if isMatch {
				// Priority
				/*if value, _ := params["priority"]; value != "" {
					if val.Priority != value {
						isMatch = false
					}
				}*/
				if isMatch {
					// Status
					if value, _ := params["status"]; value != "" {
						if val.Status != value {
							isMatch = false
						}
					}
					if isMatch {
						// Domain
						if value, _ := params["domain"]; value != "" {
							if val.Domain != value {
								isMatch = false
							}
						}
						if isMatch {
							// Asset ID
							if value, _ := params["asset_id"]; value != "" {
								assetID, _ := uuid.FromString(value)
								if *val.AssetID != assetID {
									isMatch = false
								}
							}
							if isMatch {
								// Category
								if value, _ := params["category"]; value != "" {
									if val.Category != value {
										isMatch = false
									}
								}

								/*if isMatch {
									// Due Start Date & Due End Date
									start, _ := params["due_start"]
									end, _ := params["due_end"]

									if (start != "") && (end != "") {
										start_date, err := time.Parse(time.RFC3339Nano, start)

										if err == nil {
											end_date, err := time.Parse(time.RFC3339Nano, end)

											if err == nil {
												if !checkWithinTimeRange(start_date, end_date, *val.DueDate) {
													is_match = false
												}
											}
										}
									}
								}*/

							}
						}
					}
				}
			}
			if isMatch {
				devices = append(devices, val)
			}
		}

		result <- query.Result{Result: devices}

		close(result)
	}()

	return result
}

// CountAll is used
func (r DeviceReadQueryInMemory) CountAll() <-chan query.Result {
  result := make(chan query.Result)

  go func() {
    r.Storage.Lock.RLock()
    defer r.Storage.Lock.RUnlock()

    total := len(r.Storage.DeviceReadMap)

    result <- query.Result{Result: total}

    close(result)
  }()

  return result
}

// CountDevicesWithFilter is used ...
func (r DeviceReadQueryInMemory) CountDevicesWithFilter(params map[string]string) <-chan query.Result {
  result := make(chan query.Result)

  go func() {
    r.Storage.Lock.RLock()
    defer r.Storage.Lock.RUnlock()

    devices := []storage.DeviceRead{}
    for _, val := range r.Storage.DeviceReadMap {
      isMatch := true
      // Is Due
     /* if value, _ := params["is_due"]; value != "" {
        b, _ := strconv.ParseBool(value)
        if val.IsDue != b {
          isMatch = false
        }
      }*/
      if isMatch {
        // Priority
        /*if value, _ := params["priority"]; value != "" {
          if val.Priority != value {
            isMatch = false
          }
		}*/
        if isMatch {
          // Status
          if value, _ := params["status"]; value != "" {
            if val.Status != value {
              isMatch = false
			}
          }
          if isMatch {
            // Domain
            if value, _ := params["domain"]; value != "" {
              if val.Domain != value {
                isMatch = false
              }
            }
            if isMatch {
              // Asset ID
              if value, _ := params["asset_id"]; value != "" {
                assetID, _ := uuid.FromString(value)
                if *val.AssetID != assetID {
                  isMatch = false
                }
              }
              if isMatch {
                // Category
                if value, _ := params["category"]; value != "" {
                  if val.Category != value {
                    isMatch = false
                  }
                }
                /*if isMatch {
                  // Due Start Date & Due End Date
                  start, _ := params["due_start"]
                  end, _ := params["due_end"]

                  if (start != "") && (end != "") {
                    start_date, err := time.Parse(time.RFC3339Nano, start)

                    if err == nil {
                      end_date, err := time.Parse(time.RFC3339Nano, end)

                      if err == nil {
                        if !checkWithinTimeRange(start_date, end_date, *val.DueDate) {
                          is_match = false
                        }
                      }
                    }
                  }
                }*/
              }
            }
          }
        }
      }
      if isMatch {
        devices = append(devices, val)
      }
    }

    result <- query.Result{Result: len(devices)}

    close(result)
  }()

  return result
}

/*
func checkWithinTimeRange(start time.Time, end time.Time, check time.Time) bool {

	is_start := check.Equal(start)
	is_end := check.Equal(end)
	is_between := check.After(start) && check.Before(end)
	return is_start || is_end || is_between
}*/