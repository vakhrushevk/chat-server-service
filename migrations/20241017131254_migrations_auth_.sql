-- +goose Up
create table chat
(
    id         serial primary key,
    name       text not null,
    created_at timestamp,
    updated_at timestamp
);

create table chat_user
(
    id         serial primary key,
    id_chat    integer references chat (id) on delete cascade,
    id_user    integer,
    created_at timestamp,
    updated_at timestamp
);

create table messages
(
    id         serial primary key,
    sender     integer,
    text       text,
    id_chat    integer references chat (id) on delete cascade,
    created_at timestamp,
    updated_at timestamp

);

-- +goose Down
drop table chat_user;
drop table messages;
drop table chat;
