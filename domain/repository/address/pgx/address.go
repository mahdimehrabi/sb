package pgx

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"m1-article-service/domain/entity"
	"m1-article-service/infrastructure/godotenv"
)

const pageSize = 10

type AddressRepository struct {
	env  *godotenv.Env
	conn *pgxpool.Pool
}

func NewAddressRepository(env *godotenv.Env, conn *pgxpool.Pool) *AddressRepository {
	lr := &AddressRepository{
		env:  env,
		conn: conn,
	}
	return lr
}

func (r AddressRepository) Create(ctx context.Context, address *entity.Address) (int64, error) {
	sql := `INSERT INTO address (address_id, street, city, state, zip_code, country) VALUES($1, $2, $3, $4, $5, $6) RETURNING address_id`
	err := r.conn.
		QueryRow(ctx, sql,
			address.ID, address.Street, address.City, address.State, address.ZipCode, address.Country).Scan(&address.ID)
	if err != nil {
		return 0, err
	}
	return address.ID, nil
}

func (r AddressRepository) Update(ctx context.Context, address *entity.Address) error {
	sql := `UPDATE address SET street=$1, city=$2, state=$3, zip_code=$4, country=$5 WHERE address_id=$6`
	if _, err := r.conn.Exec(ctx, sql,
		address.Street, address.City, address.State, address.ZipCode, address.Country, address.ID); err != nil {
		return err
	}
	return nil
}

func (r AddressRepository) Delete(ctx context.Context, addressID int64) error {
	sql := `DELETE FROM address WHERE address_id=$1`
	if _, err := r.conn.Exec(ctx, sql, addressID); err != nil {
		return err
	}
	return nil
}

func (r AddressRepository) Detail(ctx context.Context, addressID int64) (address *entity.Address, err error) {
	address = new(entity.Address)
	sql := `SELECT address_id, street, city, state, zip_code, country FROM address WHERE address_id=$1`
	err = r.conn.QueryRow(ctx, sql, addressID).
		Scan(&address.ID, &address.Street, &address.City, &address.State, &address.ZipCode, &address.Country)
	if err != nil {
		return nil, err
	}
	return
}

func (r AddressRepository) List(ctx context.Context, pageNumber uint16, pageSize uint16) ([]*entity.Address, error) {
	addresses := make([]*entity.Address, 0)
	offset := (pageNumber - 1) * pageSize
	sql := `SELECT address_id, street, city, state, zip_code, country FROM address LIMIT $1 OFFSET $2`
	rows, err := r.conn.Query(ctx, sql, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		address := &entity.Address{}
		if err := rows.Scan(&address.ID, &address.Street, &address.City, &address.State, &address.ZipCode, &address.Country); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return addresses, nil
}
