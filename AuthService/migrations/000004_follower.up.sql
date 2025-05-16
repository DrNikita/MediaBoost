CREATE TABLE IF NOT EXISTS follower (
    id SERIAL PRIMARY KEY,
    user_id BIGINT references user(id),
    followed_by BIGINT references user(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CHECK (user_id != followed_by),
    UNIQUE(user_id, followed_by)
);

CREATE INDEX idx_follower_user_id ON follower(user_id);
CREATE INDEX idx_follower_followed_by ON follower(followed_by);