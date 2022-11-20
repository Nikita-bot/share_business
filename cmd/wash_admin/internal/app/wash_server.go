package app

import (
	"context"
	"wash_admin/internal/entity"
	"wash_admin/internal/entity/vo"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type WashServerService interface {
	GetWashServer(ctx context.Context, auth *Auth, id uuid.UUID) (entity.WashServer, error)
	AddWashServer(ctx context.Context, auth *Auth, addWashServer vo.AddWashServer) error
}

type Repository interface {
	GetWashAdmin(ctx context.Context, identity string) (entity.WashAdmin, error)
	GetWashServer(ctx context.Context, ownerId uuid.UUID, id uuid.UUID) (entity.WashServer, error)
	AddWashServer(ctx context.Context, addWashServer vo.AddWashServer, ownerId uuid.UUID) error
}

type WashServerSvc struct {
	l    *zap.SugaredLogger
	repo Repository
}

func NewWashServerService(logger *zap.SugaredLogger, repo Repository) WashServerService {
	return &WashServerSvc{
		l:    logger,
		repo: repo,
	}
}

func (wa *WashServerSvc) GetWashServer(ctx context.Context, auth *Auth, id uuid.UUID) (entity.WashServer, error) {
	owner, err := wa.repo.GetWashAdmin(ctx, auth.UID)

	if err != nil {
		return entity.WashServer{}, err
	}

	return wa.repo.GetWashServer(ctx, owner.ID, id)
}

func (wa *WashServerSvc) AddWashServer(ctx context.Context, auth *Auth, addWashServer vo.AddWashServer) error {
	owner, err := wa.repo.GetWashAdmin(ctx, auth.UID)

	if err != nil {
		return err
	}

	return wa.repo.AddWashServer(ctx, addWashServer, owner.ID)
}
