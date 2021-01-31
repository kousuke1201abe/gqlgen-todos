CREATE TABLE IF NOT EXISTS todos (
  id varchar(64) PRIMARY KEY,
  text varchar(256) NOT NULL,
  done bool NOT NULL,
  user_id varchar(64) NOT NULL,
  FOREIGN KEY fk_todos_1(user_id) REFERENCES users(id)
);
