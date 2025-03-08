package lsp

import "sync"

type State struct {
	documents map[string]documentState
	mutex     sync.RWMutex
}

type documentState struct {
	content string
	version int
}

func NewState() *State {
	return &State{documents: make(map[string]documentState)}
}

func (s *State) read(uri string) string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.documents[uri].content
}

func (s *State) update(uri, content string, version int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if document, ok := s.documents[uri]; ok && document.version > version {
		return
	}
	s.documents[uri] = documentState{content: content, version: version}
}

func (s *State) close(uri string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.documents, uri)
}
