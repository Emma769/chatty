CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE function update_timestamp() returns TRIGGER as $$
begin
  new.updated_at = now();
  return new;
end;
$$ language plpgsql;

CREATE OR REPLACE function inc_version() returns TRIGGER as $$
begin
  new.version = new.version + 1;
  return new;
end;
$$ language plpgsql;
