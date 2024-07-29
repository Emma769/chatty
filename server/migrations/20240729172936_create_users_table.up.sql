CREATE TABLE IF NOT EXISTS users(
  user_id TEXT default replace(uuid_generate_v4()::text, '-', ''),
  username TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password BYTEA NOT NULL,
  version INT default 1,
  created_at TIMESTAMP WITH TIME ZONE default now(), 
  updated_at TIMESTAMP WITH TIME ZONE,
  deleted_at TIMESTAMP WITH TIME ZONE,
  PRIMARY KEY(user_id)
);
