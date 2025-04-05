CREATE TABLE IF NOT EXISTS users
(
    id       uuid primary key ,
    username text not null,
    email    text unique not null ,
    token    text not null
);

