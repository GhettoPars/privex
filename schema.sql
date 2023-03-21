CREATE TABLE messages (
  message_id BIGSERIAL PRIMARY KEY,
  user_id integer NOT NULL,
  message_text text NOT NULL,
  message_type text NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE users (
  user_id BIGSERIAL PRIMARY KEY,
  user_name text NOT NULL,
  user_role text NOT NULL,
  email text NOT NULL,
  password text NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);