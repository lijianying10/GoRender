package rt

import "sync"

// DataChangeEvent a component
type DataChangeEvent struct {
	channels     map[string]chan struct{}
	channelsLock sync.RWMutex
}

// NewDataChangeEvent new DataChangeEvent
func NewDataChangeEvent() *DataChangeEvent {
	return &DataChangeEvent{
		channels: make(map[string]chan struct{}),
	}
}

// Bind bind a new client
func (bcust *DataChangeEvent) Bind(id string) chan struct{} {
	bcust.channelsLock.Lock()
	defer bcust.channelsLock.Unlock()
	bcust.channels[id] = make(chan struct{}, 10)
	return bcust.channels[id]
}

// Unbind a client
func (bcust *DataChangeEvent) Unbind(id string) {
	bcust.channelsLock.Lock()
	defer bcust.channelsLock.Unlock()
	close(bcust.channels[id])
	delete(bcust.channels, id)
}

// BCust a event
func (bcust *DataChangeEvent) BCust(event struct{}) {
	bcust.channelsLock.RLock()
	defer bcust.channelsLock.RUnlock()
	for _, channel := range bcust.channels {
		channel <- event
	}
}
