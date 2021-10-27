package repo

import pb "bitbucket.org/udevs/sharqtv_go_user_service/genproto/user_service"

type RoleRepoI interface {
	Create(role *pb.Role) (string, error)
}
