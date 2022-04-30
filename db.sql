CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE customers(
  id BIGSERIAL PRIMARY KEY
  name VARCHAR(50) NOT NULL
  code VARCHAR(50) NOT NULL DEFAULT uuid_generate_v4 ()
);

CREATE TABLE orders(
  id BIGSERIAL PRIMARY KEY
  item VARCHAR(100) NOT NULL
  amount INTEGER
  user_id BIGINT REFERENCES "customers" ("id")
  created_at TIMESTAMPZ DEFAULT (now() at time zone 'utc')
)