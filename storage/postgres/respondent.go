package postgres

import (
	"fmt"

	pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"

	"bitbucket.org/udevs/ur_go_user_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type respondentRepo struct {
	db *sqlx.DB
}

// NewRespondentRepo ...
func NewRespondentRepo(db *sqlx.DB) repo.RespondentRepoI {
	return &respondentRepo{db: db}
}

func (r *respondentRepo) Create(respondent *pb.Respondent) (string, error) {
	query := `INSERT INTO respondent(
                            id,
                            name,
                            email,
                            phone,
                            sber_id)
                    VALUES ($1, $2, $3, $4, $5) `

	_, err := r.db.Exec(
		query,
		respondent.Id,
		respondent.Name,
		respondent.Email,
		respondent.Phone,
		respondent.SberId,
	)

	return respondent.Id, err
}

func (r *respondentRepo) Update(respondent *pb.Respondent) (string, error) {
	query := `UPDATE respondent 
                    SET
                        name = $1,
                        email = $2,
                        phone = $3,
                        sber_id = $4,
                        updated_at = current_timestamp
                WHERE id = $5 `

	_, err := r.db.Exec(
		query,
		respondent.Name,
		respondent.Email,
		respondent.Phone,
		respondent.SberId,
		respondent.Id,
	)

	return respondent.Id, err
}

func (r *respondentRepo) Get(id string) (*pb.Respondent, error) {
	var respondent pb.Respondent
	query := `SELECT 
                    id,
                    name,
                    email,
                    phone,
                    sber_id
                FROM respondent
                WHERE deleted_at = 0 AND id = $1 `

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&respondent.Id,
		&respondent.Name,
		&respondent.Email,
		&respondent.Phone,
		&respondent.SberId,
	)
	if err != nil {
		return nil, err
	}

	return &respondent, nil
}

func (r *respondentRepo) GetAll(req *pb.GetAllRespondentRequest) (*pb.GetAllRespondentResponse, error) {
	var (
		filter      string
		count       int32
		respondents []*pb.Respondent
		args        = make(map[string]interface{})
	)

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

	countQuery := `SELECT count(1) FROM respondent WHERE deleted_at = 0 ` + filter
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

	filter += " ORDER BY created_at DESC LIMIT :limit OFFSET :offset "
	args["limit"] = req.Limit
	args["offset"] = req.Offset

	query := `SELECT
                    id,
                    name,
                    email,
                    phone,
                    sber_id
                FROM respondent 
                WHERE deleted_at = 0 %s`
	rows, err = r.db.NamedQuery(fmt.Sprintf(query, filter), args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var respondent pb.Respondent
		err = rows.Scan(
			&respondent.Id,
			&respondent.Name,
			&respondent.Email,
			&respondent.Phone,
			&respondent.SberId,
		)
		if err != nil {
			return nil, err
		}

		respondents = append(respondents, &respondent)
	}

	return &pb.GetAllRespondentResponse{
		Respondents: respondents,
		Count:       count,
	}, nil
}

func (r *respondentRepo) Delete(id string) error {
	query := `UPDATE respondent 
                SET 
                    deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int 
                WHERE id = $1 AND deleted_at=0 `

	_, err := r.db.Exec(query, id)
	return err
}
