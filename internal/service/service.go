package service

import (
	gofoundpb "gofound-grpc/api/gen/v1"
	"gofound-grpc/internal/searcher"
)

type GofoundService struct {
	Container *searcher.Container
	gofoundpb.UnimplementedGofoundServiceServer
}

// NewGofoundService 初始化服务
func NewGofoundService() *GofoundService {
	return &GofoundService{}
}
