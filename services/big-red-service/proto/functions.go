package proto

import (
	big_red_service "big-red/services/big-red-service/proto/big-red-service"
	"context"
)

type Server struct {
	big_red_service.UnimplementedBigRedServiceServer
}

func (s *Server) GetBigRed(
	ctx context.Context,
	r *big_red_service.GetBigRedRequest,
) (*big_red_service.GetBigRedResponse, error) {
	return &big_red_service.GetBigRedResponse{Name: "yo"}, nil
}
