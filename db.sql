CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS customers(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  code VARCHAR(50) NOT NULL DEFAULT uuid_generate_v4 ()
);
CREATE TABLE IF NOT EXISTS orders(
  id BIGSERIAL PRIMARY KEY,
  item VARCHAR(100) NOT NULL,
  amount INTEGER,
  customer_id BIGINT REFERENCES "customers" ("id"),
  created_at TIMESTAMPTZ DEFAULT (now() at time zone 'utc')
);
INSERT INTO customers(name) 
VALUES
('edwin'),
('walela');
INSERT INTO orders (item,amount,customer_id)
VALUES
('burger',500,1),
('chips',200,2);