package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	// "fmt"
)

type SessionService interface {
	GetSessionByEmail(email string) (model.Session, error)
}

type sessionService struct {
	sessionRepo repo.SessionRepository
}

func NewSessionService(sessionRepo repo.SessionRepository) *sessionService {
	return &sessionService{sessionRepo}
}

func (c *sessionService) GetSessionByEmail(email string) (model.Session, error) {
	session, err := c.sessionRepo.SessionAvailEmail(email)
	if err != nil {
		return model.Session{}, err
	}

	if c.sessionRepo.TokenExpired(session) {
		err := c.sessionRepo.DeleteSession(session.Token)
		if err != nil {
			return model.Session{}, err
		}
	return model.Session{}, nil // TODO: replace this
	}

	return session, nil
}