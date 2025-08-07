CREATE TABLE IF NOT EXISTS users(
  id bigserial PRIMARY KEY,
  username varchar(255) UNIQUE NOT NULL,
  password bytea NOT NULL
);
