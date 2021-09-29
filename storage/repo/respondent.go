package repo

import pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"

type RespondentRepoI interface {
	Create(respondent *pb.CreateRespondent) (string, error)
	Update(respondent *pb.UpdateRespondent) (string, error)
	Get(id string) (*pb.Respondent, error)
	GetAll(req *pb.GetAllRespondentRequest) (*pb.GetAllRespondentResponse, error)
	Delete(id string) error
	UpdatePhoto(user_id, photo string) error
	UpdateRating(rating *pb.UpdateRespondentRating) error
	UpdateRespondentInn(req *pb.UpdateRespondentInnRequest) error
	GetRespondentsById(req *pb.GetRespondentsByIdRequest) (*pb.GetAllRespondentResponse, error)
}
