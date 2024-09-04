package services

import (
	"context"
)

type ProtectedService interface {
    Admin(ctx context.Context) (string, error)
}

type protectedService struct {
}

func NewProtectedservice() ProtectedService {
    return &protectedService{}
}

func (p *protectedService) Admin(ctx context.Context) (string, error) {
    session := "session_id"
    return session, nil
}
