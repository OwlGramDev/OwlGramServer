package emoji

import "sync"

type memorySession struct {
	mux  sync.RWMutex
	data []byte
}
