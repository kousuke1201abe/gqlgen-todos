CREATE TABLE IF NOT EXISTS users (
  id varchar(64) PRIMARY KEY,
  name varchar(256) UNIQUE NOT NULL,
  index idx_users_name (name)
);
