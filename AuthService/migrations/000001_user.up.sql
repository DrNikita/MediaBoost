CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    user_role INT REFERENCES roles(id),
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    nickname VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    second_name VARCHAR(50) NOT NULL,
    age INT CHECK (age > 0),
    img_path VARCHAR(255),
    vip_expiration_date TIMESTAMPTZ NOT NULL DEFAULT NOW()
);