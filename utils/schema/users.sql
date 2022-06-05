CREATE TABLE public.users
(
    user_id VARCHAR(20) NOT NULL PRIMARY KEY,
    user_name VARCHAR(30) NOT NULL,
    full_name VARCHAR(255) NOT NULL DEFAULT '',
    password VARCHAR(255) NOT NULL DEFAULT '',
    is_admin bool NOT NULL DEFAULT false,
    create_time timestamp not null default now(),
    update_time timestamp
);

CREATE UNIQUE INDEX user_is_admin ON users (user_id, is_admin);