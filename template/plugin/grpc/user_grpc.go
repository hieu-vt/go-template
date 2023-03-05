package appgrpc

import (
	"context"
	"flag"
	"github.com/200Lab-Education/go-sdk/logger"
	"go-template/common"
	"go-template/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userClient struct {
	prefix      string
	url         string
	gwSupported bool
	gwPort      int
	client      user.UserServiceClient
}

func NewUserClient(prefix string) *userClient {
	return &userClient{
		prefix: prefix,
	}
}

func (uc *userClient) GetPrefix() string {
	return uc.prefix
}

func (uc *userClient) Get() interface{} {
	return uc
}

func (uc *userClient) Name() string {
	return uc.prefix
}

func (uc *userClient) InitFlags() {
	flag.StringVar(&uc.url, uc.GetPrefix()+"-url", "localhost:50051", "URL connect to grpc server")
}

func (uc *userClient) Configure() error {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	cc, err := grpc.Dial(uc.url, opts)

	if err != nil {
		return err
	}

	uc.client = user.NewUserServiceClient(cc)

	return nil
}

func (uc *userClient) Run() error {
	return uc.Configure()
}

func (uc *userClient) Stop() <-chan bool {
	c := make(chan bool)

	go func() {
		c <- true
	}()
	return c
}

func (uc *userClient) GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error) {
	logger.GetCurrent().GetLogger(uc.prefix).Infoln("GetUsers grpc store running")

	userIds := make([]int32, len(ids))

	for i := range userIds {
		userIds[i] = int32(ids[i])
	}

	rs, err := uc.client.GetUserByIds(ctx, &user.UserRequest{UserIds: userIds})

	if err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(rs.Users))

	for i, item := range rs.Users {

		users[i] = common.SimpleUser{
			SqlModel: common.SqlModel{
				Id: int(item.Id),
			},
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Role:      item.Role,
		}

		users[i].Mask()
	}

	return users, nil
}
