package storage

import (
	"github.com/shibukawa/configdir"
	"sync"
)

// New creates a new StorageCollection.
func New(name string) *StorageCollection {
	return &StorageCollection{
		stores: make(map[string]*Store),
		lock:   &sync.RWMutex{},
		Name:   name,
	}
}

// StorageCollection contains all of the storage interfaces that we need in the
// application. You should only need one StorageCollection in your application.
type StorageCollection struct {
	stores map[string]*Store
	lock   *sync.RWMutex
	Name   string
}

// Store is one part of a larger StorageCollection. We can have as many Stores
// as we need. When creating a Store, it should be added into the main
// StorageCollection.
type Store struct {
	cdir  configdir.ConfigDir
	main  *configdir.Config
	cache *configdir.Config
}

// GetStore returns the store for the folder/directory given.
func (sc *StorageCollection) GetStore(folder string) *Store {
	sc.lock.RLock()
	defer sc.lock.RUnlock()
	return sc.stores[folder]
}

// NewStorage creates a new store and adds it to the StorageCollection and
// returns it. If the folder/directory store already exists in the collection,
// nothing will be created, but it will return the existing Store.
func (sc *StorageCollection) NewStorage(folder string) *Store {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	// If it already exists, do nothing, but if it does not exist yet, create
	// the new store folder and add it to our list of stores.
	if _, exists := sc.stores[folder]; !exists {
		store := Store{}
		store.cdir = configdir.New(sc.Name, folder)
		store.main = store.cdir.QueryFolders(configdir.Global)[0]
		store.cache = store.cdir.QueryCacheFolder()

		sc.stores[folder] = &store
	}

	return sc.stores[folder]
}

// GetPath returns the path of the directory of the Store.
func (s *Store) GetPath() string {
	return s.main.Path
}

// GetCachePath returns the cache path of the directory of the Store. Cache can
// be deleted at any time on the OS, and should only be used for actual caching
// objects instead of important data.
func (s *Store) GetCachePath() string {
	return s.cache.Path
}
