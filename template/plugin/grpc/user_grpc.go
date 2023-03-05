package appgrpc

import (
	"context"
	"flag"
	"github.com/200Lab-Education/go-sdk/logger"
	common2 "go-template/template/common"
	user2 "go-template/template/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userClient struct {
	prefix      string
	url         string
	gwSupported bool
	gwPort      int
	client      user2.UserServiceClient
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

	uc.client = user2.NewUserServiceClient(cc)

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

func (uc *userClient) GetUsers(ctx context.Context, ids []int) ([]common2.SimpleUser, error) {
	logger.GetCurrent().GetLogger(uc.prefix).Infoln("GetUsers grpc store running")

	userIds := make([]int32, len(ids))

	for i := range userIds {
		userIds[i] = int32(ids[i])
	}

	rs, err := uc.client.GetUserByIds(ctx, &user2.UserRequest{UserIds: userIds})

	if err != nil {
		return nil, common2.ErrDB(err)
	}

	users := make([]common2.SimpleUser, len(rs.Users))

	for i, item := range rs.Users {

		users[i] = common2.SimpleUser{
			SqlModel: common2.SqlModel{
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
