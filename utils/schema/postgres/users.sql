CREATE TABLE IF NOT EXISTS users
(
    user_id     VARCHAR(20)  NOT NULL PRIMARY KEY,
    user_name   VARCHAR(30)  NOT NULL,
    full_name   VARCHAR(255) NOT NULL DEFAULT '',
    password    VARCHAR(255) NOT NULL DEFAULT '',
    is_admin    bool         NOT NULL DEFAULT false,
    create_time timestamp    not null default now(),
    update_time timestamp
);

CREATE UNIQUE INDEX IF NOT EXISTS user_is_admin ON users (user_id, is_admin);

-- password : admintokobaba
INSERT INTO users (user_id, user_name, full_name, password, is_admin, create_time, update_time)
VALUES ('admin', 'admin', 'admin', '$2a$10$N2jeiBRc/licIaVfC.wZmuG72x3yDF.GxdES8L0hDBm6Jla2IWeTC', true,
        '2022-07-22 02:36:31.189328', null)
ON CONFLICT DO NOTHING;
