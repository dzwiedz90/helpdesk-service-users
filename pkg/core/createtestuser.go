package core

import (
	"context"

	"github.com/dzwiedz90/helpdesk-service-users/model"
	"github.com/dzwiedz90/helpdesk-service-users/service/serviceconfig"
)

func NewTestCore() TestCore {
	return TestCore{}
}

type TestCore struct{}

func (c *TestCore) CreateUser(ctx context.Context, cfg *serviceconfig.ServerConfig, creq *model.CreateUser) (int64, error) {
	return 0, nil
}

func (c *TestCore) GetUser(ctx context.Context, cfg *serviceconfig.ServerConfig, id int64) (*model.User, error) {
	return nil, nil
}

func (c *TestCore) GetAlUsers(ctx context.Context, cfg *serviceconfig.ServerConfig) ([]*model.User, error) {
	return nil, nil
}

func (c *TestCore) UpdateUser(ctx context.Context, cfg *serviceconfig.ServerConfig, creq *model.UpdateUser) error {
	return nil
}

func (c *TestCore) DeleteUser(ctx context.Context, cfg *serviceconfig.ServerConfig, id int) error {
	return nil
}
