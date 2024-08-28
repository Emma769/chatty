CREATE TABLE IF NOT EXISTS sessions (
  scope TEXT NOT NULL check(scope = 'authentication'),
  hash BYTEA NOT NULL,
  valid_till TIMESTAMP WITH TIME ZONE,
  is_revoked BOOLEAN DEFAULT FALSE,
  user_id TEXT NOT NULL,
  email TEXT NOT NULL,
  CONSTRAINT sessions_users_fk FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE,
  CONSTRAINT sessions_users_fk_2 FOREIGN KEY(email) REFERENCES users(email)
);

CREATE INDEX sessions_hash_idx ON sessions(hash);

CREATE OR REPLACE function remove_old_session_rows() returns TRIGGER AS $$
BEGIN
  DELETE FROM sessions WHERE valid_till < now() + interval '1 min';
  return new;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER remove_old_session_rows_trigger AFTER INSERT ON sessions EXECUTE PROCEDURE remove_old_session_rows();
