CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE books
(
    id          serial       not null unique,
    title       varchar(255) not null,
    author      varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_books
(
    id      serial                                      not null unique,
    user_id int references users (id) on delete cascade not null,
    book_id int references books (id) on delete cascade not null
);