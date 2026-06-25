package battery

import (
	"sync"
	"time"
)

type Level struct {
	DeviceID  string    `json:"device_id"`
	Percent   int       `json:"percent"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Service store and retrieve battery levels in memory
type Service struct {
	mu     sync.RWMutex
	levels map[string]*Level
}

// NewService creates a new battery service
func NewService() *Service {
	return &Service{
		levels: make(map[string]*Level),
	}
}

// Update sets the battery level for a device
func (s *Service) Update(deviceID string, percent int) *Level {
	s.mu.Lock()
	defer s.mu.Unlock()
	l := &Level{
		DeviceID:  deviceID,
		Percent:   percent,
		UpdatedAt: time.Now(),
	}
	s.levels[deviceID] = l
	return l
}

func (s *Service) Get(deviceID string) (*Level, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	l, ok := s.levels[deviceID]
	return l, ok
}

// Clear battery level for a device (e.g on disconnect).
func (s *Service) Remove(deviceID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.levels, deviceID)
}

func (s *Service) All() []*Level {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*Level, 0, len(s.levels))
	for _, l := range s.levels {
		list = append(list, l)
	}
	return list
}
