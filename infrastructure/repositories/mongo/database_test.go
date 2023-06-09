package mongo_test

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MongoDBTestContainer struct {
	testcontainers.Container
}

func NewMongoDBTestContainer(ctx context.Context) (*MongoDBTestContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "mongo:6.0.5",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor: wait.ForAll(
			wait.ForLog("Waiting for connections"),
			wait.ForListeningPort("27017/tcp"),
		),
	}

	container, err := testcontainers.GenericContainer(
		ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	if err != nil {
		return nil, err
	}

	return &MongoDBTestContainer{
		Container: container,
	}, nil
}
