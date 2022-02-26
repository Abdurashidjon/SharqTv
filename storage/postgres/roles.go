package postgres

import (
	//"database/sql"
	"database/sql"

	pb "bitbucket.org/udevs/sharqtv_go_user_service/genproto/user_service"
	"bitbucket.org/udevs/sharqtv_go_user_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type roleRepo struct {
	db *sqlx.DB
}

// NewRespondentRepo ...
func NewRoleRepo(db *sqlx.DB) repo.RoleRepoI {
	return &roleRepo{db: db}
}
func (r *roleRepo) Create(roles *pb.Role) (string, error) {
	query := `INSERT INTO admin_roles(
                            id,
                            title,
                            control_movie_creators,
                            control_category,
                            control_favorite_movie,
                            control_dashboard,
                            control_tariff,
                            control_permissions,
                            control_movie,
                            control_notification,
                            control_admin,
                            control_genre )
                    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) `
	_, err := r.db.Exec(
		query,
		roles.Id,
		roles.Title,
		roles.ControlMovieCreators,
		roles.ControlCategory,
		roles.ControlFavoriteMovie,
		roles.ControlDashboard,
		roles.ControlTariff,
		roles.ControlPermissions,
		roles.ControlMovie,
		roles.ControlNotification,
		roles.ControlAdmin,
		roles.ControlGenre,
	)

	if err != nil {
		return "", err
	}
	return roles.Id, err
}

func (r *roleRepo) GetRoleID(role *pb.GetRequest) (*pb.GetResponse, error) {
	var (
		resp pb.GetResponse
	)

	query := `SELECT 
	                     id,
	                     title,
	                     control_movie_creators,
	                     control_category,
	                     control_favorite_movie,
	                     control_dashboard,
	                     control_tariff,
	                     control_permissions,
	                     control_movie,
	                     control_notification,
	                     control_admin,
	                    control_genre
                FROM admin_roles
                WHERE deleted_at is null AND id = $1 `

	row := r.db.QueryRow(query, role.GetId())
	err := row.Scan(
		&resp.Id,
		&resp.Title,
		&resp.ControlMovieCreators,
		&resp.ControlCategory,
		&resp.ControlFavoriteMovie,
		&resp.ControlDashboard,
		&resp.ControlTariff,
		&resp.ControlPermissions,
		&resp.ControlMovie,
		&resp.ControlNotification,
		&resp.ControlAdmin,
		&resp.ControlGenre,
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}


func (r *roleRepo) DeleteRoleId(id string) error {

	result, err := r.db.Exec(`DELETE FROM admin_roles WHERE id=$1`, id)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}




func (r *roleRepo) UpdateRole(role *pb.UpdateList) (*pb.UpdateResponse, error) {

 	query := `UPDATE admin_roles
	                SET
	                title = $1,
	                control_movie_creators = $2,
	                control_category = $3,
	                control_favorite_movie = $4,
	                control_dashboard = $5,
	                control_tariff = $6,
	                control_permissions = $7,
	                control_movie = $8,
	                control_notification = $9,
	                control_admin = $10,
	                control_genre = $11,
	                updated_at = CURRENT_TIMESTAMP
	                WHERE id = $12 AND active=true`

	_, err := r.db.Exec(
		query,
		role.Title,
		role.ControlMovieCreators,
		role.ControlCategory,
		role.ControlFavoriteMovie,
		role.ControlDashboard,
		role.ControlTariff,
		role.ControlPermissions,
		role.ControlMovie,
		role.ControlNotification,
		role.ControlNotification,
		role.ControlAdmin,
		role.ControlGenre,
	)
	if err != nil {
		return (*pb.UpdateResponse)(role), err
	}

	return (*pb.UpdateResponse)(role), nil 
}



/* func (r *roleRepo) ListRole(req *pb.ListRequest) (*pb.RoleListResponse, error){
	var (
		filter string
		count int32
		roles []*pb.RoleList
		args = make.(map[string]interface{})
	)
} */
