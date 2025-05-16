CREATE TABLE IF NOT EXISTS "role" (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL
);

INSERT INTO "role"(role_name) VALUES ('master'), ('slave');