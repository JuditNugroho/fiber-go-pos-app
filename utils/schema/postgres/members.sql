CREATE TABLE IF NOT EXISTS members
(
    id VARCHAR(20) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    create_time timestamp not null default now(),
    update_time timestamp
);


CREATE INDEX IF NOT EXISTS member_phone_idx ON members (phone);