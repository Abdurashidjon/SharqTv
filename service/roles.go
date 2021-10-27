package service

import (
	"context"

	pb "bitbucket.org/udevs/sharqtv_go_user_service/genproto/user_service"
	"bitbucket.org/udevs/sharqtv_go_user_service/pkg/logger"
	"bitbucket.org/udevs/sharqtv_go_user_service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type rolesService struct {
	logger  logger.Logger
	storage storage.StorageI
}

func NewRoleService(db *sqlx.DB, log logger.Logger) *rolesService {
	return &rolesService{
		logger:  log,
		storage: storage.NewStoragePg(db),
	}
}

func (s *rolesService) Create(ctx context.Context, req *pb.Role) (*pb.IdMessage, error) {
	id, err := s.storage.Roles().Create(req)
	if err != nil {
		s.logger.Error("Error while creating Admin in service", logger.Error(err), logger.Any("req", req))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.IdMessage{
		Id: id,
	}, nil
}
