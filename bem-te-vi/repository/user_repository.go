package repository

import (
	"application/domain"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user domain.User) error {

	query := `INSERT INTO users (name, email, password, document) VALUES ($1, $2, $3, $4) RETURNING id`

	return ur.db.QueryRow(context.Background(), query, user.Name, user.Email, user.Password, user.Document).Scan(&user.ID)
}

func (ur *UserRepository) FindById(id int) (domain.User, error) {

	var user domain.User

	query := `SELECT id, name, email FROM users WHERE id = $1`

	err := ur.db.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email)

	return user, err
}

func (ur *UserRepository) FindByEmail(username string) (domain.User, error) {

	var user domain.User

	query := `SELECT id, name, email, password FROM users WHERE email = $1`

	err := ur.db.QueryRow(context.Background(), query, username).Scan(&user.ID, &user.Name, &user.Email)

	return user, err
}

func (ur *UserRepository) FindAll() (pgx.Rows, error) {
	query := `SELECT id, name, email, document FROM users`

	return ur.db.Query(context.Background(), query)

}

func (ur *UserRepository) Update(user domain.User) (pgconn.CommandTag, error) {
	query := `UPDATE users SET name = $1 WHERE id = $2`

	return ur.db.Exec(context.Background(), query, user.Name, user.ID)
}

func (ur *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := ur.db.Exec(context.Background(), query, id)
	return err
}
