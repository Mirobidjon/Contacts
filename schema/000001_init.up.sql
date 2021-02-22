CREATE TABLE users (
    id          serial          not null unique,
    name        varchar(30)     not null,
    username    varchar(30)     not null unique,
    password    varchar(255)    not null
);

CREATE TABLE contacts (
    id serial not null unique,
    name varchar(30) not null,
    phone varchar(30) not null,
    user_id int references users (id) on delete cascade not null
);