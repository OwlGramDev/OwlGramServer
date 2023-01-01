package emoji

import (
	"OwlGramServer/consts"
	"context"
	"os"
)

func (s *memorySession) StoreSession(_ context.Context, data []byte) error {
	s.mux.Lock()
	s.data = data
	go func() {
		_ = os.WriteFile(consts.TelegramSessionFile, s.data, 0644)
	}()
	s.mux.Unlock()
	return nil
}
