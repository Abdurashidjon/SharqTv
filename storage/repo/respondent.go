package repo

import pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"

type RespondentRepoI interface {
	Create(respondent *pb.Respondent) (string, error)
	Update(respondent *pb.Respondent) (string, error)
	Get(id string) (*pb.Respondent, error)
	GetAll(req *pb.GetAllRespondentRequest) (*pb.GetAllRespondentResponse, error)
	Delete(id string) error
}
