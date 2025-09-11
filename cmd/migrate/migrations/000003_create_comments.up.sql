CREATE TABLE IF NOT EXISTS comments(
  id bigserial PRIMARY KEY,
  user_id bigint REFERENCES users(id),
  post_id bigint REFERENCES posts(id),
  content text NOT NULL,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  likes bigint
);
