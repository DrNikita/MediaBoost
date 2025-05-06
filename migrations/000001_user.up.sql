CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    pass VARCHAR(255) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    second_name VARCHAR(50) NOT NULL,
    nickname VARCHAR(255) UNIQUE,
    age INT CHECK (age >= 0),
    vip BOOLEAN DEFAULT FALSE,
    vip_expiration_date DATE
);