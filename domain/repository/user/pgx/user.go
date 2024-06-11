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
	sql := `INSERT INTO users (firstname, lastname) VALUES($1, $2) RETURNING id`
	err := r.conn.QueryRow(ctx, sql, user.Name, user.Lastname).Scan(&user.ID)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r UserRepository) Detail(ctx context.Context, userID int64) (user *entity.User, err error) {
	user = new(entity.User)
	sql := `SELECT id, firstname, lastname FROM users WHERE id=$1`
	err = r.conn.QueryRow(ctx, sql, userID).Scan(&user.ID, &user.Name, &user.Lastname)
	if err != nil {
		return nil, err
	}
	return user, nil
}
