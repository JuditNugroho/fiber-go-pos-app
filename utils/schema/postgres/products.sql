CREATE TABLE IF NOT EXISTS products
(
    product_id VARCHAR(20) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    barcode VARCHAR(30) NOT NULL DEFAULT '0',
    stock BIGINT NOT NULL DEFAULT 0,
    ppn BOOLEAN NOT NULL DEFAULT FALSE ,
    price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    member_price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    discount NUMERIC(10, 2) NOT NULL DEFAULT 0,
    category_id INTEGER NOT NULL DEFAULT 0,
    create_time timestamp not null default now(),
    update_time timestamp
);


CREATE INDEX IF NOT EXISTS barcode_products_idx ON products (barcode);