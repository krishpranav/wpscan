package database

import (
	"sync"
)

var (
	Memory = NewMemory()
	mutex  sync.Mutex
)

type memory struct {
	stringx   map[string]string
	intx      map[string]int
	slice     map[string][]string
	boolx     map[string]bool
	mapstring map[string]map[string]string
}

func NewMemory() *memory {
	database := &memory{
		stringx:   map[string]string{},
		intx:      map[string]int{},
		boolx:     map[string]bool{},
		slice:     map[string][]string{},
		mapstring: map[string]map[string]string{},
	}

	database.mapstring["HTTP Plugins Versions"] = map[string]string{}
	database.mapstring["HTTP Themes Versions"] = map[string]string{}

	return database
}

func (db *memory) SetString(key, value string) {
	mutex.Lock()
	db.stringx[key] = value
	mutex.Unlock()
}

func (db *memory) SetSlice(key string, value []string) {
	mutex.Lock()
	db.slice[key] = value
	mutex.Unlock()
}

func (db *memory) SetInt(key string, value int) {
	mutex.Lock()
	db.intx[key] = value
	mutex.Unlock()
}

func (db *memory) SetBool(key string, value bool) {
	mutex.Lock()
	db.boolx[key] = value
	mutex.Unlock()
}
