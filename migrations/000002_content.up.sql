CREATE TABLE IF NOT EXISTS content (
    id SERIAL PRIMARY KEY,
    user_id BIGINT references user(id),
    content_path VARCHAR
);