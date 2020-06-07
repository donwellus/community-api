package topic

import (
	"community-api/repositories"
	"sync"
)

// InMemory represents a memory service for Topics
type InMemory struct {
	data map[repositories.TopicCode]repositories.Topic
	mux  sync.RWMutex
}

// NewTopicInMemoryRepository instantiate a new Repository for Topic
func NewTopicInMemoryRepository() *InMemory {
	return &InMemory{
		data: map[repositories.TopicCode]repositories.Topic{
			"xpto": {
				Code: "xpto",
				Name: "XPTO",
			},
			"otpx": {
				Code: "otpx",
				Name: "OTPX",
			},
			"internet_speed": {
				Code: "internet_speed",
				Name: "Internet Speed",
			},
		},
	}
}

// List all Topics
func (r InMemory) List() (list []repositories.Topic) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	for _, t := range r.data {
		list = append(list, t)
	}
	return list
}

// Get a Topic by code
func (r InMemory) Get(code repositories.TopicCode) (*repositories.Topic, error) {
	r.mux.RLock()
	found, ok := r.data[code]
	r.mux.RUnlock()

	if !ok {
		return nil, nil //TODO: return error
	}

	return &found, nil
}
