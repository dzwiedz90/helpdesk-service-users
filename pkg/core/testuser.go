package core

import (
	"context"
	"errors"

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
	if id > 0 {
		return &model.User{
			UserId:    1,
			Username:  "james.kirk",
			Email:     "james.kirk@enterprise.com",
			FirstName: "James",
			LastName:  "Kirk",
			Age:       32,
			Gender:    "male",
			Address: model.Address{
				Street:     "Kirksuckz 8",
				City:       "Deep Space Nine",
				PostalCode: "2137CU666",
				Country:    "Federation",
			},
		}, nil
	} else {
		return nil, errors.New("sql: no rows in result set")
	}
}

func (c *TestCore) GetAlUsers(ctx context.Context, cfg *serviceconfig.ServerConfig) ([]*model.User, error) {
	return []*model.User{
		{
			UserId:    1,
			Username:  "james.kirk",
			Email:     "james.kirk@enterprise.com",
			FirstName: "James",
			LastName:  "Kirk",
			Age:       32,
			Gender:    "male",
			Address: model.Address{
				Street:     "Spocksuckz 8",
				City:       "Deep Space Nine",
				PostalCode: "2137CU666",
				Country:    "Federation",
			},
		},
		{
			UserId:    1,
			Username:  "mr.spock",
			Email:     "mr.spock@enterprise.com",
			FirstName: "Mr",
			LastName:  "Spock",
			Age:       97,
			Gender:    "male",
			Address: model.Address{
				Street:     "Kirksuckz 8",
				City:       "Deep Space Nine",
				PostalCode: "2137CU666",
				Country:    "Federation",
			},
		},
	}, nil
}

func (c *TestCore) UpdateUser(ctx context.Context, cfg *serviceconfig.ServerConfig, creq *model.UpdateUser) error {
	return nil
}

func (c *TestCore) DeleteUser(ctx context.Context, cfg *serviceconfig.ServerConfig, id int64) error {
	return nil
}
