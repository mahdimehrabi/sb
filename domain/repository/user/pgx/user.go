package pgx

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"m1-article-service/domain/entity"
	userRepo "m1-article-service/domain/repository/user"
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

func (r UserRepository) Detail(ctx context.Context, userID int64) (*entity.User, error) {
	sql := `
		SELECT 
			u.id, u.firstname, u.lastname, 
			a.id, a.city, a.state, a.country, a.street, a.zip_code, a.user_id
		FROM 
			users u
		LEFT JOIN 
			addresses a ON u.id = a.user_id
		WHERE 
			u.id = $1
	`

	rows, err := r.conn.Query(ctx, sql, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := &entity.User{
		Addresses: []*entity.Address{},
	}

	for rows.Next() {
		var (
			address entity.Address
		)

		err := rows.Scan(
			&user.ID, &user.Name, &user.Lastname,
			&address.ID, &address.City, &address.State, &address.Country, &address.Street, &address.ZipCode, &address.UserID,
		)
		if err != nil {
			return nil, err
		}

		// Append address if it's not null
		if address.ID != 0 {
			user.Addresses = append(user.Addresses, &address)
		}
	}

	if rows.Err() != nil {
		if errors.Is(rows.Err(), pgx.ErrNoRows) {
			return nil, userRepo.ErrNotFound
		}
		return nil, rows.Err()
	}
	if user.ID == 0 {
		return nil, userRepo.ErrNotFound
	}

	return user, nil
}
