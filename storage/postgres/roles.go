package postgres

import (
	pb "bitbucket.org/udevs/sharqtv_go_user_service/genproto/user_service"

	"bitbucket.org/udevs/sharqtv_go_user_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type roleRepo struct {
	db *sqlx.DB
}

// NewRespondentRepo ...
func NewRoleRepo(db *sqlx.DB) repo.RoleRepoI {
	return &roleRepo{
		db: db,
	}
}

func (r *roleRepo) Create(roles *pb.Role) (string, error) {
	query := `INSERT INTO roles(
                            id,
                            title,
                            control_movie_creators
							control_category
							control_favorite_movie
							control_dashboard
							control_tariff
							control_permissions
							control_movie
							control_notification
							control_admin
							control_genre
							)
                    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12,) `
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
