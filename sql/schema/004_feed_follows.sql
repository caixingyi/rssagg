-- +goose up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,    -- 主键，用 UUID 作为唯一标识
    created_at TIMESTAMP NOT NULL,    -- 创建时间
    updated_at TIMESTAMP NOT NULL,    -- 更新时间
    user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,    -- 用户 ID
    feed_id UUID REFERENCES feeds(id) ON DELETE CASCADE NOT NULL,    -- 订阅源 ID
    UNIQUE(user_id, feed_id)    -- 确保每个用户只能关注同一个订阅源一次
);


-- +goose down
DROP TABLE feed_follows;
