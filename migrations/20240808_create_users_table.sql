-- +goose Up
create table users (
    chat_id int references chats(id) on delete cascade,
    user_id int not null,
    primary key (chat_id, user_id)
);

-- +goose Down
drop table users;