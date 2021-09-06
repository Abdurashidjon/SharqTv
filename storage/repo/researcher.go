package repo

import pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"

type ResearcherRepoI interface {
	Create(researcher *pb.Researcher) (string, error)
	Update(researcher *pb.Researcher) (string, error)
	Get(id string) (*pb.Researcher, error)
	GetAll(req *pb.GetAllResearcherRequest) (*pb.GetAllResearcherResponse, error)
	Delete(id string) error
	UpdatePhoto(user_id, photo string) error
}
