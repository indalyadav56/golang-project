create table if not exists users(
    id int primary key generated always as identity,
    full_name varchar(255),
    email text,
    password text,
    created_at timestamp default now()
);

create table if not exists todos(
    id int primary key generated always as identity,
    user_id int references users(id),
    title text,
    created_at timestamp default now(),
    completed_at timestamp,
    updated_at timestamp default now(),
    deleted_at timestamp default now()
);
