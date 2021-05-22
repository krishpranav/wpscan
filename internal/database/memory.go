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
