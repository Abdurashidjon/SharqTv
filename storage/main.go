package storage

import (
	"bitbucket.org/udevs/sharqtv_go_user_service/storage/postgres"
	"bitbucket.org/udevs/sharqtv_go_user_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Roles() repo.RoleRepoI
}

type storagePg struct {
	db       *sqlx.DB
	roleRepo repo.RoleRepoI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:       db,
		roleRepo: postgres.NewRoleRepo(db),
	}
}

func (s *storagePg) Roles() repo.RoleRepoI {
	return s.roleRepo
}

