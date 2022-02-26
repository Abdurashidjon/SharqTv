package service

import (
	"context"
	"database/sql"
	"fmt"

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
	pb.UnimplementedRolesServiceServer
}

func NewRoleService(db *sqlx.DB, log logger.Logger) *rolesService {
	return &rolesService{
		logger:  log,
		storage: storage.NewStoragePg(db),
	}
}

func (s *rolesService) Create(ctx context.Context, req *pb.Role) (*pb.IdMessage, error) {
	id, err := s.storage.Roles().Create(req)
	fmt.Println("function called")
	if err != nil {
		s.logger.Error("Error while creating Admin in service", logger.Error(err), logger.Any("req", req))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.IdMessage{
		Id: id,
	}, nil
}

func (s *rolesService) RoleGetID(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	resp, err := s.storage.Roles().GetRoleID(req)
	if err != nil {
		s.logger.Error("Error while getID role in service", logger.Error(err), logger.Any("req", req))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *rolesService) DeleteRoleId(ctx context.Context, req *pb.GetRequest) (*pb.Empty, error) {
	err := s.storage.Roles().DeleteRoleId(req.GetId())
	if err == sql.ErrNoRows {
		s.logger.Error("Error while deleting role,Not found", logger.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not Found")
	} else if err != nil {
		s.logger.Error("Error while deleting Role", logger.Error(err), logger.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &pb.Empty{}, nil
}

func (s *rolesService) UpdateRole(ctx context.Context, req *pb.UpdateList) (*pb.UpdateResponse, error) {

	role, err := s.storage.Roles().UpdateRole(req)
	if err == sql.ErrNoRows {
		s.logger.Error("Error whiling updating role not found", logger.Any("req", req))
	} else if err != nil {
		s.logger.Error("Error while updating Role", logger.Error(err), logger.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server errror")
	}
	return &pb.UpdateResponse{
		Id:                   role.Id,
		Title:                role.Title,
		ControlCategory:      role.ControlCategory,
		ControlMovieCreators: role.ControlMovieCreators,
		ControlFavoriteMovie: role.ControlFavoriteMovie,
		ControlDashboard:     role.ControlDashboard,
		ControlTariff:        role.ControlTariff,
		ControlPermissions:   role.ControlPermissions,
		ControlMovie:         role.ControlMovie,
		ControlNotification:  role.ControlNotification,
		ControlAdmin:         role.ControlAdmin,
		ControlGenre:         role.ControlGenre,
	}, nil

	/* 	id, err := s.storage.Roles().UpdateRole(req)
	   	if err != nil {
	   		return nil, helper.HandleError(s.logger, err, "Error whiling updating new role", req, codes.Internal)
	   	}
	   	return &pb.UpdateResponse{
	   		Id: id,
	   	}, nil */

	/* resp, err := s.storage.Roles().UpdateRole(req)
	if err != nil {
		s.logger.Error("Error while UpdateId role in service", logger.Error(err), logger.Any("req", req))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil */
}

// func (s *rolesService) ListRole(ctx context.Context, req *pb.ListRequest) (*pb.RoleListResponse, error) {
// 	resp, err := s.storage.Roles().ListRole(req)
// 	if err != nil {
// 		return nil, helper.HandleError(s.logger, err, "error while getting all roles", req, codes.Internal)
// 	}
// 	return resp, nil  }
