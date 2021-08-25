package repo

import pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"

type CompanyRepoI interface {
	Create(company *pb.Company) (string, error)
	Update(company *pb.Company) (string, error)
	Get(id string) (*pb.Company, error)
	GetAll(req *pb.GetAllCompanyRequest) (*pb.GetAllCompanyResponse, error)
	Delete(id string) error
}