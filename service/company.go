package service

import (
	"context"
	"errors"
	"fmt"

	pbb "bitbucket.org/udevs/ur_go_user_service/genproto/billing_service"
	pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"
	"bitbucket.org/udevs/ur_go_user_service/pkg/helper"
	"bitbucket.org/udevs/ur_go_user_service/pkg/logger"
	"bitbucket.org/udevs/ur_go_user_service/pkg/util"
	grpc_client "bitbucket.org/udevs/ur_go_user_service/service/grpc_clients"
	"bitbucket.org/udevs/ur_go_user_service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type companyService struct {
	logger  logger.Logger
	storage storage.StorageI
	client  *grpc_client.GrpcClients
}

func NewCompanyService(db *sqlx.DB, log logger.Logger, transaction *grpc_client.GrpcClients) *companyService {
	return &companyService{
		logger:  log,
		storage: storage.NewStoragePg(db),
		client:  transaction,
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

	account_number, err := s.client.AccountService().Create(
		context.Background(),
		&pbb.CreateAccountReq{
			Accounts: []*pbb.CreateAccount{
				{
					UserId:        req.Id,
					Name:          req.Name,
					AccountTypeId: "company",
				},
			},
		},
	)

	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating account", req, codes.Internal)
	}

	req.AccountNumber = account_number.AccountNumber
	fmt.Println(account_number.AccountNumber)
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
