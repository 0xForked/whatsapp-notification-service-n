package service

import (
	"context"
	"github.com/aasumitro/gowa/domain"
)

type gowansService struct {
	ctx               context.Context
	sessionRepository domain.ISessionRepository
}

func NewGowansService(
	ctx context.Context,
	sessionRepository domain.ISessionRepository,
) domain.IGowansService {
	return &gowansService{
		ctx:               ctx,
		sessionRepository: sessionRepository,
	}
}
