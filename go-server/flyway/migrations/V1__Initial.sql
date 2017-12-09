CREATE TABLE users (
  user_id      BIGSERIAL PRIMARY KEY,
  steam_id     VARCHAR(17),
  display_name VARCHAR(64),
  session_id   VARCHAR(50)
);

CREATE INDEX ON users (session_id);
CREATE INDEX ON users (steam_id);