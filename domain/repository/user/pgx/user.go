package pgx

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"m1-article-service/domain/entity"
	"m1-article-service/infrastructure/godotenv"
)

const pageSize = 10

type UserRepository struct {
	env  *godotenv.Env
	conn *pgxpool.Pool
}

func NewUserRepository(env *godotenv.Env, conn *pgxpool.Pool) *UserRepository {
	ur := &UserRepository{
		env:  env,
		conn: conn,
	}
	return ur
}

func (r UserRepository) Create(ctx context.Context, user *entity.User) (int64, error) {
	sql := `INSERT INTO user_user (name, lastname) VALUES($1, $2) RETURNING id`
	err := r.conn.QueryRow(ctx, sql, user.Name, user.Lastname).Scan(&user.ID)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r UserRepository) Update(ctx context.Context, user *entity.User) error {
	sql := `UPDATE user_user SET name=$1, lastname=$2 WHERE id=$3`
	if _, err := r.conn.Exec(ctx, sql, user.Name, user.Lastname, user.ID); err != nil {
		return err
	}
	return nil
}

func (r UserRepository) Delete(ctx context.Context, userID int64) error {
	sql := `DELETE FROM user_user WHERE id=$1`
	if _, err := r.conn.Exec(ctx, sql, userID); err != nil {
		return err
	}
	return nil
}

func (r UserRepository) Detail(ctx context.Context, userID int64) (user *entity.User, err error) {
	user = new(entity.User)
	sql := `SELECT id, name, lastname FROM user_user WHERE id=$1`
	err = r.conn.QueryRow(ctx, sql, userID).Scan(&user.ID, &user.Name, &user.Lastname)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r UserRepository) List(ctx context.Context, pageNumber uint16, pageSize uint16) ([]*entity.User, error) {
	users := make([]*entity.User, 0)
	offset := (pageNumber - 1) * pageSize
	sql := `SELECT id, name, lastname FROM user_user LIMIT $1 OFFSET $2`
	rows, err := r.conn.Query(ctx, sql, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &entity.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Lastname); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return users, nil
}
