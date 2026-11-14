-- +goose Up

CREATE TABLE feeds (
    id UUID PRIMARY KEY,    -- 主键，用 UUID 作为唯一标识
    created_at TIMESTAMP NOT NULL,    -- 创建时间
    updated_at TIMESTAMP NOT NULL,    -- 更新时间
    name TEXT NOT NULL,               -- 订阅源名称
    url TEXT UNIQUE NOT NULL,   -- 订阅源 URL
    user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL    -- 用户 ID
);

-- +goose Down
DROP TABLE feeds;