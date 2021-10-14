package postgres

import (
	"fmt"

	pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"

	"bitbucket.org/udevs/ur_go_user_service/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type respondentRepo struct {
	db *sqlx.DB
}

// NewRespondentRepo ...
func NewRespondentRepo(db *sqlx.DB) repo.RespondentRepoI {
	return &respondentRepo{db: db}
}

func (r *respondentRepo) Create(respondent *pb.CreateRespondent) (string, error) {
	query := `INSERT INTO respondent(
                            id,
                            name,
                            email,
                            phone,
                            sber_id,
                            company,
                            position,
                            description,
							account_number
							)
                    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) `
	_, err := r.db.Exec(
		query,
		respondent.Id,
		respondent.Name,
		respondent.Email,
		respondent.Phone,
		respondent.SberId,
		respondent.Company,
		respondent.Position,
		respondent.Description,
		respondent.AccountNumber,
	)

	if err != nil {
		return "", err
	}
	return respondent.Id, err
}

func (r *respondentRepo) Update(respondent *pb.UpdateRespondent) (string, error) {
	query := `UPDATE respondent 
                    SET
                        name = $1,
                        email = $2,
                        phone = $3,
                        sber_id = $4,
                        company = $5,
                        position = $6,
                        description = $7,
                        updated_at = current_timestamp
                WHERE id = $8 `

	_, err := r.db.Exec(
		query,
		respondent.Name,
		respondent.Email,
		respondent.Phone,
		respondent.SberId,
		respondent.Company,
		respondent.Position,
		respondent.Description,
		respondent.Id,
	)
	if err != nil {
		return "", err
	}

	return respondent.Id, err
}

func (r *respondentRepo) Get(id string) (*pb.Respondent, error) {
	var respondent pb.Respondent
	var rating pb.Rating
	query := `SELECT 
                    id,
                    name,
                    email,
                    phone,
                    sber_id,
					inn,
                    company,
                    position,
                    description,
                    photo,
					account_number,
					rating_communication,
					rating_experience,
					rating_punctuality
                FROM respondent
                WHERE deleted_at = 0 AND id = $1 `

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&respondent.Id,
		&respondent.Name,
		&respondent.Email,
		&respondent.Phone,
		&respondent.SberId,
		&respondent.Inn,
		&respondent.Company,
		&respondent.Position,
		&respondent.Description,
		&respondent.Photo,
		&respondent.AccountNumber,
		&rating.Communication,
		&rating.Experience,
		&rating.Punctuality,
	)
	if err != nil {
		return nil, err
	}
	respondent.Rating = &rating
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

	if req.AccountNumber == 1 {
		filter += " AND account_number = 0 "
	} else if req.AccountNumber > 0 {
		filter += " AND account_number = :account_number "
		args["account_number"] = req.AccountNumber
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
                    sber_id,
					inn,
                    company,
                    position,
                    description,
                    photo,
					account_number,
					rating_communication,
					rating_experience,
					rating_punctuality
                FROM respondent 
                WHERE deleted_at = 0 %s`
	rows, err = r.db.NamedQuery(fmt.Sprintf(query, filter), args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var respondent pb.Respondent
		var rating pb.Rating
		err = rows.Scan(
			&respondent.Id,
			&respondent.Name,
			&respondent.Email,
			&respondent.Phone,
			&respondent.SberId,
			&respondent.Inn,
			&respondent.Company,
			&respondent.Position,
			&respondent.Description,
			&respondent.Photo,
			&respondent.AccountNumber,
			&rating.Communication,
			&rating.Experience,
			&rating.Punctuality,
		)
		if err != nil {
			return nil, err
		}
		respondent.Rating = &rating

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

func (r *respondentRepo) UpdatePhoto(user_id, photo string) error {
	query := `UPDATE respondent
                SET
                    photo = $1,
                    updated_at = current_timestamp
                WHERE id = $2 `

	_, err := r.db.Exec(query, photo, user_id)
	return err
}

func (r *respondentRepo) UpdateRating(rating *pb.UpdateRespondentRating) error {
	query := `UPDATE respondent
				SET
					rating_communication = $1,
					rating_experience = $2,
					rating_punctuality = $3
				WHERE 
					id = $4`
	_, err := r.db.Exec(
		query,
		rating.Communication,
		rating.Experience,
		rating.Punctuality,
		rating.UserId,
	)
	return err
}

func (r *respondentRepo) UpdateRespondentInn(req *pb.UpdateRespondentInnRequest) error {
	query := `UPDATE respondent
				SET
					inn = $1
				WHERE 
					id = $2`
	_, err := r.db.Exec(
		query,
		req.Inn,
		req.RespondentId,
	)

	fmt.Println("respondent inn" + req.Inn)
	return err
}

func (r *respondentRepo) GetRespondentsById(req *pb.GetRespondentsByIdRequest) (*pb.GetAllRespondentResponse, error) {
	var (
		filter      string
		respondents []*pb.Respondent
		args        = make(map[string]interface{})
	)
	args["ids"] = pq.Array(req.Ids)

	filter += " ORDER BY created_at DESC LIMIT :limit OFFSET :offset "
	args["limit"] = req.Limit
	args["offset"] = req.Offset

	query := `SELECT
                    id,
                    name,
                    email,
                    phone,
                    sber_id,
					inn,
                    company,
                    position,
                    description,
                    photo,
					account_number,
					rating_communication,
					rating_experience,
					rating_punctuality
                FROM respondent 
                WHERE deleted_at = 0 AND id = ANY (:ids) %s`
	rows, err := r.db.NamedQuery(fmt.Sprintf(query, filter), args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var respondent pb.Respondent
		var rating pb.Rating
		err = rows.Scan(
			&respondent.Id,
			&respondent.Name,
			&respondent.Email,
			&respondent.Phone,
			&respondent.SberId,
			&respondent.Inn,
			&respondent.Company,
			&respondent.Position,
			&respondent.Description,
			&respondent.Photo,
			&respondent.AccountNumber,
			&rating.Communication,
			&rating.Experience,
			&rating.Punctuality,
		)
		if err != nil {
			return nil, err
		}
		respondent.Rating = &rating

		respondents = append(respondents, &respondent)
	}

	return &pb.GetAllRespondentResponse{
		Respondents: respondents,
	}, nil
}

func (r *respondentRepo) UpdateAccountNumber(req *pb.CreateRespondent) error {
	query := `UPDATE company
				SET
					account_number = $1
				WHERE id = $2`

	_, err := r.db.Exec(query, req.AccountNumber, req.Id)
	return err
}
