package server

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type ShareNotFoundError struct {
	ID string
}

func (e *ShareNotFoundError) Error() string {
	return fmt.Sprintf("share [%s] not found", e.ID)
}

type Share struct {
	ContentType string
	Content     []byte
}

type ShareStore struct {
	ShareMap map[string]Share
	mu       sync.RWMutex
}

func NewShareStore() *ShareStore {
	return &ShareStore{
		ShareMap: make(map[string]Share),
		mu:       sync.RWMutex{},
	}
}

func (ss *ShareStore) AddShare(contentType string, content []byte) (string, error) {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	uuidString := id.String()

	ss.ShareMap[uuidString] = Share{
		contentType,
		content,
	}

	return uuidString, nil
}

func (ss *ShareStore) GetShare(id string) (Share, error) {
	ss.mu.RLock()
	defer ss.mu.RUnlock()

	share, ok := ss.ShareMap[id]
	if !ok {
		return Share{}, &ShareNotFoundError{id}
	}

	return share, nil
}

func (ss *ShareStore) DeleteShare(id string) {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	if _, ok := ss.ShareMap[id]; ok {
		delete(ss.ShareMap, id)
	}
}
