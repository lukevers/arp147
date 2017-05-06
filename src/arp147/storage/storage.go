package storage

import (
	"github.com/lukevers/storage"
)

// Collection contains all of the directories for temporary and permanent
// storage on the disk.
var Collection *storage.StorageCollection

func init() {
	Collection = storage.New("arp147")
}
