package manager

/*
import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

type userSession struct {
	name    string
	timeout int
}

// creates sessions and keeps track of users
type userManager struct {
	logger   *slog.Logger
	sessions map[userSession]bool
}

func (m *userManager) Login(name string, timeout int) {
	newSession = userSession{name: name, timeout: timeout}
	m.sessions[newSession] = true
}

func (m *userManager) Logout(session userSession) {
	delete(m.sessions, session)
	close(session.jobChan)
}
*/
