package session

import (
	"errors"
	"sync"
	"time"
)

type manager struct {
	idToSession  map[string]*session
	usernameToId map[string]string
	lock         sync.Mutex
}

func NewManager() *manager {
	return &manager{
		idToSession:  make(map[string]*session),
		usernameToId: make(map[string]string),
	}
}

// returns username, isManager
// error if the session no exists
func (m *manager) GetInfo(id string) (string, bool, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	s, ok := m.idToSession[id]
	if !ok {
		return "", false, errors.New("no session " + id)
	}
	return s.username, s.isManager, nil
}

func (m *manager) IsIdActive(id string) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	_, ok := m.idToSession[id]
	return ok
}

func (m *manager) IsUsernameActive(username string) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	_, ok := m.usernameToId[username]
	return ok
}

func (m *manager) AddSession(username string, isManager bool) (string, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	_, ok := m.usernameToId[username]
	if ok {
		return "", errors.New("username is logged")
	}
	s := newSession(username, isManager)
	m.usernameToId[username] = s.id
	m.idToSession[s.id] = s
	return s.id, nil
}

func (m *manager) DeleteSession(id string) error {
	if !m.IsIdActive(id) {
		return errors.New("no session to delete")
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	s, _ := m.idToSession[id]
	delete(m.usernameToId, s.username)
	delete(m.idToSession, id)
	return nil
}

func (m *manager) UpdateSession(id string) error {
	if !m.IsIdActive(id) {
		return errors.New("no session " + id)
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	s, _ := m.idToSession[id]
	s.expire = time.Now().Add(time.Hour) //
	return nil
}
