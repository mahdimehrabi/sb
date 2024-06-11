CREATE TABLE IF NOT EXISTS users (
                                    id BIGSERIAL PRIMARY KEY,
                                    firstname varchar(50) NOT NULL,
                                    lastname varchar(50) NOT NULL
);


CREATE TABLE IF NOT EXISTS addresses (
                                         id BIGSERIAL PRIMARY KEY,
                                         street VARCHAR(255),
    city VARCHAR(255),
    state VARCHAR(255),
    zip_code VARCHAR(10),
    country VARCHAR(255),
    user_id  BIGINT REFERENCES users(id)
);