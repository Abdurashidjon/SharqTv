package service

import (
	"context"
	"fmt"

	pbb "bitbucket.org/udevs/ur_go_user_service/genproto/billing_service"
	pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"
	"bitbucket.org/udevs/ur_go_user_service/pkg/helper"
	"bitbucket.org/udevs/ur_go_user_service/pkg/logger"
	grpc_client "bitbucket.org/udevs/ur_go_user_service/service/grpc_clients"
	"bitbucket.org/udevs/ur_go_user_service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type respondentService struct {
	logger  logger.Logger
	storage storage.StorageI
	client  *grpc_client.GrpcClients
}

func NewRespondentService(db *sqlx.DB, log logger.Logger, transaction *grpc_client.GrpcClients) *respondentService {
	return &respondentService{
		logger:  log,
		storage: storage.NewStoragePg(db),
		client:  transaction,
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
	account_number, err := s.client.AccountService().Create(
		context.Background(),
		&pbb.CreateAccountReq{
			Accounts: []*pbb.CreateAccount{
				{
					UserId:        req.Id,
					Name:          req.Name,
					AccountTypeId: "respondent",
				},
			},
		},
	)

	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating account", req, codes.Internal)
	}

	req.AccountNumber = account_number.AccountNumber
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

func (s *respondentService) UpdateAccountNumber(ctx context.Context, req *pb.CreateRespondent) (*emptypb.Empty, error) {
	err := s.storage.Respondent().UpdateAccountNumber(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating respondent account number", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}
