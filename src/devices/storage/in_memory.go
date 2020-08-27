package storage

import (
	"fmt"
	"time"

	deadlock "github.com/sasha-s/go-deadlock"
	uuid "github.com/satori/go.uuid"
)

// DeviceEventStorage is used ...
type DeviceEventStorage struct {
	Lock       *deadlock.RWMutex
	DeviceEvents []DeviceEvent
}

// CreateDeviceEventStorage is used
func CreateDeviceEventStorage() *DeviceEventStorage {
	rwMutex := deadlock.RWMutex{}
	deadlock.Opts.DeadlockTimeout = time.Second * 10
	deadlock.Opts.OnPotentialDeadlock = func() {
		fmt.Println("Device EVENT STORAGE DEADLOCK!")
	}

	return &DeviceEventStorage{Lock: &rwMutex}
}

// DeviceReadStorage is used
type DeviceReadStorage struct {
	Lock        *deadlock.RWMutex
	DeviceReadMap map[uuid.UUID]DeviceRead
}

// CreateDeviceReadStorage is used ...
func CreateDeviceReadStorage() *DeviceReadStorage {
	rwMutex := deadlock.RWMutex{}
	deadlock.Opts.DeadlockTimeout = time.Second * 10
	deadlock.Opts.OnPotentialDeadlock = func() {
		fmt.Println("DEVICE READ STORAGE DEADLOCK!")
	}

	return &DeviceReadStorage{DeviceReadMap: make(map[uuid.UUID]DeviceRead), Lock: &rwMutex}
}
