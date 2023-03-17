CREATE TABLE messages (
  message_id   BIGSERIAL PRIMARY KEY,
  user_id integer NOT NULL,
  message_text text NOT NULL,
  message_type text NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
