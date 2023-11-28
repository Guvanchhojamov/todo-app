CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id serial not null unique,
    title varchar(255) not null,
    description text not null
);

CREATE TABLE todo_items
(
    id serial not null unique,
    title varchar(255) not null,
    description text not null,
    done boolean default false not null
);

CREATE TABLE users_lists
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null ,
    list_id int REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE lists_items
(
    id serial not null unique,
    list_id int references todo_lists(id) on delete cascade not null ,
    item_id int references todo_items(id) on DELETE cascade not null
);
