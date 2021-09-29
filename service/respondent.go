package service

import (
	"context"
	"fmt"

	pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"
	"bitbucket.org/udevs/ur_go_user_service/pkg/helper"
	"bitbucket.org/udevs/ur_go_user_service/pkg/logger"
	"bitbucket.org/udevs/ur_go_user_service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type respondentService struct {
	logger  logger.Logger
	storage storage.StorageI
}

func NewRespondentService(db *sqlx.DB, log logger.Logger) *respondentService {
	return &respondentService{
		logger:  log,
		storage: storage.NewStoragePg(db),
	}
}

func (s *respondentService) Create(ctx context.Context, req *pb.CreateRespondent) (*pb.RespondentID, error) {
	// TODO - validate inn
	// if !util.IsValidPhone(req.Phone) {
	// 	return nil, helper.HandleError(s.logger, errors.New("Invalid phone number"), "Invalid phone number", req, codes.Canceled)
	// }

	// if !util.IsValidEmail(req.Email) {
	// 	return nil, helper.HandleError(s.logger, errors.New("Invalid email"), "Invalid email", req, codes.Canceled)
	// }
	fmt.Println("1")
	id, err := s.storage.Respondent().Create(req)
	fmt.Println("2")
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating new respondent", req, codes.Internal)
	}

	return &pb.RespondentID{
		Id: id,
	}, nil
}

func (s *respondentService) Update(ctx context.Context, req *pb.UpdateRespondent) (*pb.RespondentID, error) {
	// TODO - validate inn
	// if !util.IsValidPhone(req.Phone) {
	// 	return nil, helper.HandleError(s.logger, errors.New("Invalid phone number"), "Invalid phone number", req, codes.Canceled)
	// }

	// if !util.IsValidEmail(req.Email) {
	// 	return nil, helper.HandleError(s.logger, errors.New("Invalid email"), "Invalid email", req, codes.Canceled)
	// }
	id, err := s.storage.Respondent().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating new respondent", req, codes.Internal)
	}

	return &pb.RespondentID{
		Id: id,
	}, nil
}

func (s *respondentService) Get(ctx context.Context, req *pb.RespondentID) (*pb.Respondent, error) {
	respondent, err := s.storage.Respondent().Get(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting respondent", req, codes.Internal)
	}

	return respondent, nil
}

func (s *respondentService) GetAll(ctx context.Context, req *pb.GetAllRespondentRequest) (*pb.GetAllRespondentResponse, error) {
	respondents, err := s.storage.Respondent().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all respondents", req, codes.Internal)
	}

	return respondents, nil
}

func (s *respondentService) Delete(ctx context.Context, req *pb.RespondentID) (*emptypb.Empty, error) {
	err := s.storage.Respondent().Delete(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting respondent", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (s *respondentService) UpdatePhoto(ctx context.Context, req *pb.UpdateRespondentPhoto) (*emptypb.Empty, error) {
	err := s.storage.Respondent().UpdatePhoto(req.UserId, req.Photo)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating respondent photo", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (s *respondentService) UpdateRating(ctx context.Context, req *pb.UpdateRespondentRating) (*emptypb.Empty, error) {
	err := s.storage.Respondent().UpdateRating(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating respondent rating", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (s *respondentService) UpdateRespondentInn(ctx context.Context, req *pb.UpdateRespondentInnRequest) (*emptypb.Empty, error) {
	err := s.storage.Respondent().UpdateRespondentInn(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating respondent inn", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (s *respondentService) GetRespondentsById(ctx context.Context, req *pb.GetRespondentsByIdRequest) (*pb.GetAllRespondentResponse, error) {
	respondents, err := s.storage.Respondent().GetRespondentsById(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting respondents by id", req, codes.Internal)
	}

	return respondents, nil
}
