CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL
);

INSERT INTO roles(role_name) VALUES ('master'), ('slave');