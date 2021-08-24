package storage

import (
	"bitbucket.org/udevs/ur_go_user_service/storage/postgres"
	"bitbucket.org/udevs/ur_go_user_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Company() repo.CompanyRepoI
	Respondent() repo.RespondentRepoI
}

type storagePg struct {
	db             *sqlx.DB
	companyRepo    repo.CompanyRepoI
	respondentRepo repo.RespondentRepoI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:             db,
		companyRepo:    postgres.NewCompanyRepo(db),
		respondentRepo: postgres.NewRespondentRepo(db),
	}
}

func (s *storagePg) Company() repo.CompanyRepoI {
	return s.companyRepo
}

func (s *storagePg) Respondent() repo.RespondentRepoI {
	return s.respondentRepo
}
