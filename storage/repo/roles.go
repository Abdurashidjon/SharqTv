package repo

import pb "bitbucket.org/udevs/sharqtv_go_user_service/genproto/user_service"

type RoleRepoI interface {
	Create(role *pb.Role) (string, error)
	GetRoleID(role *pb.GetRequest) (*pb.GetResponse, error)
	DeleteRoleId(id string) error
	UpdateRole(role *pb.UpdateList) (*pb.UpdateResponse,error)
	//ListRole(role *pb.ListRequest) (*pb.RoleListResponse, error)
}
