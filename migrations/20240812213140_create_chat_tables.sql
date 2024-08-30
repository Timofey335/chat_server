-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    role INT NOT NULL
);

CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    users VARCHAR[],
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    chat_id INT REFERENCES chats(id) ON DELETE CASCADE,
    text TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE user_to_chats (
    chat_id INT REFERENCES chats(id),
    user_id INT REFERENCES users(id),
    CONSTRAINT user_to_chats_pk PRIMARY KEY (chat_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users, chats, messages, user_to_chats;
-- +goose StatementEnd
