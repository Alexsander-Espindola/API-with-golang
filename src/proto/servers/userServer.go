package servers

import (
	"context"

	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
)

type postUserClient struct{}

func (c *postUserClient) PostUser(ctx context.Context, in *pb.User) (*pb.Response, error) {

}
