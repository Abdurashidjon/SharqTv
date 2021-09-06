package postgres

import (
	"fmt"

	pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"

	"bitbucket.org/udevs/ur_go_user_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type researcherRepo struct {
	db *sqlx.DB
}

// NewResearcherRepo ...
func NewResearcherRepo(db *sqlx.DB) repo.ResearcherRepoI {
	return &researcherRepo{db: db}
}

func (r *researcherRepo) Create(researcher *pb.Researcher) (string, error) {
	query := `INSERT INTO researcher(
                            id,
                            name,
                            email,
                            phone,
                            profession_title,
                            role_id,
                            company_id)
                    VALUES ($1, $2, $3, $4, $5, $6, $7) `

	_, err := r.db.Exec(
		query,
		researcher.Id,
		researcher.Name,
		researcher.Email,
		researcher.Phone,
		researcher.ProfessionTitle,
		researcher.RoleId,
		researcher.CompanyId,
	)

	return researcher.Id, err
}

func (r *researcherRepo) Update(researcher *pb.Researcher) (string, error) {
	query := `UPDATE researcher 
                    SET
                        name = $1,
                        email = $2,
                        phone = $3,
                        profession_title = $4,
                        role_id = $5,
                        updated_at = current_timestamp
                WHERE id = $7 `

	_, err := r.db.Exec(
		query,
		researcher.Name,
		researcher.Email,
		researcher.Phone,
		researcher.ProfessionTitle,
		researcher.RoleId,
		researcher.Id,
	)

	return researcher.Id, err
}

func (r *researcherRepo) Get(id string) (*pb.Researcher, error) {
	var researcher pb.Researcher
	query := `SELECT 
                    id,
                    name,
                    email,
                    phone,
                    profession_title,
                    role_id,
                    company_id,
					photo
                FROM researcher
                WHERE deleted_at = 0 AND id = $1 `

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&researcher.Id,
		&researcher.Name,
		&researcher.Email,
		&researcher.Phone,
		&researcher.ProfessionTitle,
		&researcher.RoleId,
		&researcher.CompanyId,
		&researcher.Photo,
	)
	if err != nil {
		return nil, err
	}

	return &researcher, nil
}

func (r *researcherRepo) GetAll(req *pb.GetAllResearcherRequest) (*pb.GetAllResearcherResponse, error) {
	var (
		filter      string
		count       int32
		researchers []*pb.Researcher
		args        = make(map[string]interface{})
	)

	args["company_id"] = req.CompanyId

	if req.Name != "" {
		filter += " AND name ilike '%' || :name || '%' "
		args["name"] = req.Name
	}

	if req.Email != "" {
		filter += " AND email ilike '%' || :email || '%' "
		args["email"] = req.Email
	}

	if req.Phone != "" {
		filter += " AND phone ilike '%' || :phone || '%' "
		args["phone"] = req.Phone
	}

	countQuery := `SELECT count(1) FROM researcher WHERE deleted_at = 0 AND company_id = :company_id %s`
	rows, err := r.db.NamedQuery(fmt.Sprintf(countQuery, filter), args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, err
		}
	}

	filter += " ORDER BY created_at DESC LIMIT :limit OFFSET :offset "
	args["limit"] = req.Limit
	args["offset"] = req.Offset

	query := `SELECT
                    id,
                    name,
                    email,
                    phone,
                    profession_title,
                    role_id,
                    company_id,
					photo
                FROM researcher 
                WHERE deleted_at = 0 AND company_id = :company_id %s`
	rows, err = r.db.NamedQuery(fmt.Sprintf(query, filter), args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var researcher pb.Researcher
		err = rows.Scan(
			&researcher.Id,
			&researcher.Name,
			&researcher.Email,
			&researcher.Phone,
			&researcher.ProfessionTitle,
			&researcher.RoleId,
			&researcher.CompanyId,
			&researcher.Photo,
		)
		if err != nil {
			return nil, err
		}

		researchers = append(researchers, &researcher)
	}

	return &pb.GetAllResearcherResponse{
		Researchers: researchers,
		Count:       count,
	}, nil
}

func (r *researcherRepo) Delete(id string) error {
	query := `UPDATE researcher 
                SET 
                    deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int 
                WHERE id = $1 AND deleted_at=0 `

	_, err := r.db.Exec(query, id)
	return err
}

func (r *researcherRepo) UpdatePhoto(user_id, photo string) error {
	query := `UPDATE respondent
                SET
                    photo = $1,
                    updated_at = current_timestamp
                WHERE id = $2 `

	_, err := r.db.Exec(query, photo, user_id)
	return err
}
