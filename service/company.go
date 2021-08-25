package service

import (
	"context"
	"errors"

	pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"
	"bitbucket.org/udevs/ur_go_user_service/pkg/helper"
	"bitbucket.org/udevs/ur_go_user_service/pkg/logger"
	"bitbucket.org/udevs/ur_go_user_service/pkg/util"
	"bitbucket.org/udevs/ur_go_user_service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type companyService struct {
	logger  logger.Logger
	storage storage.StorageI
}

func NewCompanyService(db *sqlx.DB, log logger.Logger) *companyService {
	return &companyService{
		logger:  log,
		storage: storage.NewStoragePg(db),
	}
}

func (s *companyService) Create(ctx context.Context, req *pb.Company) (*pb.CompanyId, error) {
	// TODO - validate inn
	if !util.IsValidPhone(req.Phone) {
		return nil, helper.HandleError(s.logger, errors.New("Invalid phone number"), "Invalid phone number", req, codes.Canceled)
	}

	if !util.IsValidEmail(req.Email) {
		return nil, helper.HandleError(s.logger, errors.New("Invalid email"), "Invalid email", req, codes.Canceled)
	}
	id, err := s.storage.Company().Create(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating new company", req, codes.Internal)
	}

	return &pb.CompanyId{
		Id: id,
	}, nil
}

func (s *companyService) Update(ctx context.Context, req *pb.Company) (*pb.CompanyId, error) {
	// TODO - validate inn
	if !util.IsValidPhone(req.Phone) {
		return nil, helper.HandleError(s.logger, errors.New("Invalid phone number"), "Invalid phone number", req, codes.Canceled)
	}

	if !util.IsValidEmail(req.Email) {
		return nil, helper.HandleError(s.logger, errors.New("Invalid email"), "Invalid email", req, codes.Canceled)
	}
	id, err := s.storage.Company().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating new company", req, codes.Internal)
	}

	return &pb.CompanyId{
		Id: id,
	}, nil
}

func (s *companyService) Get(ctx context.Context, req *pb.CompanyId) (*pb.Company, error) {
	company, err := s.storage.Company().Get(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting company", req, codes.Internal)
	}

	return company, nil
}

func (s *companyService) GetAll(ctx context.Context, req *pb.GetAllCompanyRequest) (*pb.GetAllCompanyResponse, error) {
	companies, err := s.storage.Company().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all companies", req, codes.Internal)
	}

	return companies, nil
}

func (s *companyService) Delete(ctx context.Context, req *pb.CompanyId) (*emptypb.Empty, error) {
	err := s.storage.Company().Delete(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting company", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}
