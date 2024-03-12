package storage

import (
	"github.com/CecileTalec/vehicle-server/storage/vehiclestore"
)

type Store interface {
	Vehicle() vehiclestore.Store
}
