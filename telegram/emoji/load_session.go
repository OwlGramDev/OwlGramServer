package emoji

import (
	"OwlGramServer/consts"
	"context"
	"github.com/gotd/td/session"
	"os"
)

func (s *memorySession) LoadSession(_ context.Context) ([]byte, error) {
	if s == nil {
		return nil, session.ErrNotFound
	}
	s.mux.RLock()
	defer s.mux.RUnlock()
	file, _ := os.ReadFile(consts.TelegramSessionFile)
	s.data = file
	if len(s.data) == 0 {
		return nil, session.ErrNotFound
	}
	cpy := append([]byte(nil), s.data...)
	return cpy, nil
}
