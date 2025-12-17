package main

import (
	"os"

	"github.com/AmiyoKm/todo-micro/api-gateway/gen/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	userServiceTarget = os.Getenv("USER_SERVICE_TARGET")
)

type userServiceClient struct {
	Client userpb.UserServiceClient
	conn   *grpc.ClientConn
}

func NewUserServerClient() (*userServiceClient, error) {
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(userServiceTarget, dialOptions...)
	if err != nil {
		return nil, err
	}

	return &userServiceClient{
		Client: userpb.NewUserServiceClient(conn),
		conn:   conn,
	}, nil
}
func (c *userServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
