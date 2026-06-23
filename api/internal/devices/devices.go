package devices

import (
	"sync"
	"time"
)

// Represents a connected bluetooth device
type Device struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Alias       string `json:"alias,omitempty"`
	Address     string `json:"address"`
	ConnectedAt time.Time `json:"connected_at"`
}

// Service manages connected devices
type Service struct {
	mu sync.RWMutex
	devices map[string]*Device
}

func (s *Service) Add(d *Device) {
	s.mu.Lock()
	defer s.mu.Unlock()
	d.ConnectedAt = time.Now()
	s.devices[d.ID] = d
}

func (s *Service) Remove(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.devices, id)
}

func (s *Service) Get(id string) (*Device, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	d, ok := s.devices[id]
	return d, ok
}

func (s *Service) All() []*Device {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*Device, 0, len(s.devices))
	for _, d := range s.devices {
		list = append(list, d)
	}
	return list
}

func (s *Service) setAlias(id, alias string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	d, ok := s.devices[id]
	if !ok {
		return false
	}
	d.Alias = alias
	return true
}

// Create a new device service
func NewService() *Service {
	return &Service{
		devices: make(map[string]*Device),
	}
}
