package commands

import (
	"sync"

	"github.com/plutov/google-home-k8s/pkg/dialogflow"
)

// UserSession .
type UserSession struct {
	SessionID     string
	ContextParams map[string]interface{}
}

// UserSessionManager .
type UserSessionManager struct {
	userSessions   map[string]UserSession
	userSessionsMU *sync.RWMutex
}

// NewUserSessionManager .
func NewUserSessionManager() *UserSessionManager {
	return &UserSessionManager{
		userSessions:   make(map[string]UserSession),
		userSessionsMU: &sync.RWMutex{},
	}
}

// GetUserSession .
func (m *UserSessionManager) GetUserSession(req *dialogflow.Request) *UserSession {
	m.userSessionsMU.Lock()
	defer m.userSessionsMU.Unlock()

	if session, ok := m.userSessions[req.Session]; ok {
		return &session
	}

	session := UserSession{
		SessionID:     req.Session,
		ContextParams: make(map[string]interface{}),
	}
	m.userSessions[req.Session] = session

	return &session
}

// SaveUserSession .
func (m *UserSessionManager) SaveUserSession(session UserSession) {
	m.userSessionsMU.Lock()
	defer m.userSessionsMU.Unlock()

	if _, ok := m.userSessions[session.SessionID]; ok {
		m.userSessions[session.SessionID] = session
	}
}
