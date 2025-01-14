package mysql

import (
	"context"
	"database/sql"
	"re-home/models"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Insert(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Password, user.Email)
	if err != nil {
		log.Error("Error inserting user: ", err)
		return err
	}

	return nil
}

func (r *UserRepository) Get(ctx context.Context, username, password string) (*models.User, error) {
	user := new(models.User)

	query := "SELECT id, username, password, email FROM users WHERE username = ? AND password = ?"
	row := r.db.QueryRowContext(ctx, query, username, password)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Error("Error fetching user: ", err)
		return nil, err
	}

	return user, nil
}
