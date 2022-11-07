package serviceuser

import (
	"context"
	"session-4/common/models"

	"github.com/golang/protobuf/ptypes/empty"
)

var localStorage *models.UserList

func init() {
	localStorage = new(models.UserList)
	localStorage.List = make([]*models.User, 0)
}

type UsersServer struct {
}

func (u UsersServer) Register(ctx context.Context, req *models.User) (*empty.Empty, err error) {
	localStorage.List = append(localStorage.List, req)

	return new(empty.Empty), nil
}

func (u UsersServer) List(context.Context, *empty.Empty) (*model.UserList, error) {
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	userSrv := UsersServer{}
	model.RegisterUsersServer(srv, userSrv)

	listener, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen. Err: %+v\n", err)
	}
}
