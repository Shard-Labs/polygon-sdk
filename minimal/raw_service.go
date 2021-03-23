package minimal

import (
	"context"

	"github.com/0xPolygon/minimal/minimal/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

type rawService struct {
	proto.UnimplementedRawServer

	s *Server
}

func (r *rawService) RawMsg(ctx context.Context, req *proto.RawMsgReq) (*empty.Empty, error) {
	return nil, nil
}
