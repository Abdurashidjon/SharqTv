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

type researcherService struct {
	logger  logger.Logger
	storage storage.StorageI
}

func NewResearcherService(db *sqlx.DB, log logger.Logger) *researcherService {
	return &researcherService{
		logger:  log,
		storage: storage.NewStoragePg(db),
	}
}

func (s *researcherService) Create(ctx context.Context, req *pb.Researcher) (*pb.ResearcherId, error) {
	// TODO - validate inn
	if !util.IsValidPhone(req.Phone) {
		return nil, helper.HandleError(s.logger, errors.New("Invalid phone number"), "Invalid phone number", req, codes.Canceled)
	}

	if !util.IsValidEmail(req.Email) {
		return nil, helper.HandleError(s.logger, errors.New("Invalid email"), "Invalid email", req, codes.Canceled)
	}
	id, err := s.storage.Researcher().Create(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating new researcher", req, codes.Internal)
	}

	return &pb.ResearcherId{
		Id: id,
	}, nil
}

func (s *researcherService) Update(ctx context.Context, req *pb.Researcher) (*pb.ResearcherId, error) {
	// TODO - validate inn
	if !util.IsValidPhone(req.Phone) {
		return nil, helper.HandleError(s.logger, errors.New("Invalid phone number"), "Invalid phone number", req, codes.Canceled)
	}

	if !util.IsValidEmail(req.Email) {
		return nil, helper.HandleError(s.logger, errors.New("Invalid email"), "Invalid email", req, codes.Canceled)
	}
	id, err := s.storage.Researcher().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating new researcher", req, codes.Internal)
	}

	return &pb.ResearcherId{
		Id: id,
	}, nil
}

func (s *researcherService) Get(ctx context.Context, req *pb.ResearcherId) (*pb.Researcher, error) {
	researcher, err := s.storage.Researcher().Get(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting researcher", req, codes.Internal)
	}

	return researcher, nil
}

func (s *researcherService) GetAll(ctx context.Context, req *pb.GetAllResearcherRequest) (*pb.GetAllResearcherResponse, error) {
	companies, err := s.storage.Researcher().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all respondents", req, codes.Internal)
	}

	return companies, nil
}

func (s *researcherService) Delete(ctx context.Context, req *pb.ResearcherId) (*emptypb.Empty, error) {
	err := s.storage.Researcher().Delete(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting researcher", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (s *researcherService) UpdatePhoto(ctx context.Context, req *pb.UpdateResearcherPhoto) (*emptypb.Empty, error) {
	err := s.storage.Researcher().UpdatePhoto(req.UserId, req.Photo)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating respondent photo", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}
