package postgres

import (
	"fmt"

	pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"
	"bitbucket.org/udevs/ur_go_user_service/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type companyRepo struct {
	db *sqlx.DB
}

// NewCompanyRepo ...
func NewCompanyRepo(db *sqlx.DB) repo.CompanyRepoI {
	return &companyRepo{db: db}
}

func (r *companyRepo) Create(company *pb.Company) (string, error) {
	query := `INSERT INTO company (
                            id,
                            name,
                            inn,
                            owner_name,
                            email,
                            phone,
							account_number
						)
                        VALUES ($1, $2, $3, $4, $5, $6, $7) `

	_, err := r.db.Exec(
		query,
		company.Id,
		company.Name,
		company.Inn,
		company.OwnerName,
		company.Email,
		company.Phone,
		company.AccountNumber,
	)

	return company.Id, err
}

func (r *companyRepo) Update(company *pb.Company) (string, error) {
	query := `UPDATE company 
                    SET
                        name = $1,
                        inn = $2,
                        owner_name = $3,
                        email = $4,
                        phone = $5,
                        updated_at = current_timestamp
                WHERE id = $6`

	_, err := r.db.Exec(
		query,
		company.Name,
		company.Inn,
		company.OwnerName,
		company.Email,
		company.Phone,
		company.Id,
	)
	if err != nil {
		return "", err
	}

	return company.Id, nil
}

func (r *companyRepo) Get(id string) (*pb.Company, error) {
	var company pb.Company
	query := `SELECT 
                    id,
                    name,
                    inn,
                    owner_name,
                    email,
                    phone,
					account_number
                FROM company
                WHERE deleted_at = 0 AND id = $1 `

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&company.Id,
		&company.Name,
		&company.Inn,
		&company.OwnerName,
		&company.Email,
		&company.Phone,
		&company.AccountNumber,
	)
	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (r *companyRepo) GetAll(req *pb.GetAllCompanyRequest) (*pb.GetAllCompanyResponse, error) {
	var (
		filter    string
		count     int32
		companies []*pb.Company
		args      = make(map[string]interface{})
	)

	if req.Name != "" {
		filter += " AND name ilike '%' || :name || '%' "
		args["name"] = req.Name
	}

	if req.Inn != "" {
		filter += " AND inn ilike '%' || :inn || '%' "
		args["inn"] = req.Inn
	}

	if req.AccountNumber == 1 {
		filter += " AND account_number = 0 "
	} else if req.AccountNumber > 0 {
		filter += " AND account_number = :account_number "
		args["account_number"] = req.AccountNumber
	}

	countQuery := `SELECT count(1) FROM company WHERE deleted_at = 0 ` + filter
	rows, err := r.db.NamedQuery(countQuery, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, err
		}
	}

	filter += " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	args["limit"] = req.Limit
	args["offset"] = req.Offset

	query := `SELECT
                    id,
                    name,
                    inn,
                    owner_name,
                    email,
                    phone,
					account_number
                FROM company 
                WHERE deleted_at = 0 %s`
	rows, err = r.db.NamedQuery(fmt.Sprintf(query, filter), args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var company pb.Company
		err = rows.Scan(
			&company.Id,
			&company.Name,
			&company.Inn,
			&company.OwnerName,
			&company.Email,
			&company.Phone,
			&company.AccountNumber,
		)
		if err != nil {
			return nil, err
		}

		companies = append(companies, &company)
	}

	return &pb.GetAllCompanyResponse{
		Companies: companies,
		Count:     count,
	}, nil
}

func (r *companyRepo) Delete(id string) error {
	query := `UPDATE company 
                SET 
                    deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int 
                WHERE id = $1 AND deleted_at=0 `

	_, err := r.db.Exec(query, id)
	return err
}

func (r *companyRepo) UpdateAccountNumber(req *pb.Company) error {
	query := `UPDATE company
				SET
					account_number = $1
				WHERE id = $2`

	_, err := r.db.Exec(query, req.AccountNumber, req.Id)
	return err
}

func (r *companyRepo) CompanyReport(req *pb.CompanyReportReq) (*pb.CompanyReportResp, error) {
	var (
		filter    string
		companies []*pb.CompanyReportResponse
		args      = make(map[string]interface{})
	)
	args["ids"] = pq.Array(req.Companies)

	filter += " ORDER BY c.created_at DESC "

	query := `SELECT 
					c.id,
					c.name,
					count(1) 
				FROM researcher r JOIN company c ON c.id = r.company_id 
                WHERE c.deleted_at = 0 AND c.id = ANY (:ids) GROUP BY (c.id,c.name) %s`
	rows, err := r.db.NamedQuery(fmt.Sprintf(query, filter), args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var company pb.CompanyReportResponse
		err = rows.Scan(
			&company.CompanyId,
			&company.CompanyName,
			&company.ResearchersCount,
		)
		if err != nil {
			return nil, err
		}
		companies = append(companies, &company)
	}

	return &pb.CompanyReportResp{
		Companies: companies,
	}, nil
}

func (r *companyRepo) GetCompanyByAccountNumber(req *pb.CompanyAccountNumber) (*pb.Company, error) {
	var company pb.Company
	query := `SELECT 
                    id,
                    name,
                    inn,
                    owner_name,
                    email,
                    phone,
					account_number
                FROM company
                WHERE deleted_at = 0 AND account_number = $1 `

	row := r.db.QueryRow(query, req.AccountNumber)
	err := row.Scan(
		&company.Id,
		&company.Name,
		&company.Inn,
		&company.OwnerName,
		&company.Email,
		&company.Phone,
		&company.AccountNumber,
	)
	if err != nil {
		return nil, err
	}

	return &company, nil
}
