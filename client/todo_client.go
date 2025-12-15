package main

import (
	"os"

	"github.com/AmiyoKm/todo-micro/api-gateway/gen/todopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	todoServiceTarget = os.Getenv("TODO_SERVICE_TARGET")
)

type todoServiceClient struct {
	Client todopb.TodoServiceClient
	conn   *grpc.ClientConn
}

func NewTodoServerClient() (*todoServiceClient, error) {
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(todoServiceTarget, dialOptions...)
	if err != nil {
		return nil, err
	}

	return &todoServiceClient{
		Client: todopb.NewTodoServiceClient(conn),
		conn:   conn,
	}, nil
}
func (c *todoServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
