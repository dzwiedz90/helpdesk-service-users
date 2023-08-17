package core

import (
	"context"

	"github.com/dzwiedz90/helpdesk-service-users/model"
	"github.com/dzwiedz90/helpdesk-service-users/service/serviceconfig"
)

type CoreAdapter interface {
	CreateUser(ctx context.Context, cfg *serviceconfig.ServerConfig, creq *model.CreateUser) (int64, error)
	GetUser(ctx context.Context, cfg *serviceconfig.ServerConfig, id int64) (*model.User, error)
	GetAlUsers(ctx context.Context, cfg *serviceconfig.ServerConfig) ([]*model.User, error)
	UpdateUser(ctx context.Context, cfg *serviceconfig.ServerConfig, creq *model.UpdateUser) error
	DeleteUser(ctx context.Context, cfg *serviceconfig.ServerConfig, id int) error
}
