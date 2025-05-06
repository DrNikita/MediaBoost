CREATE TABLE IF NOT EXISTS post (
    id SERIAL PRIMARY KEY,
    user_id BIGINT references user(id),
    content_id BIGINT references content(id),
);