ALTER TABLE users
DROP COLUMN IF EXISTS active,
DROP COLUMN IF EXISTS deleted_at,
DROP COLUMN IF EXISTS updated_at;

ALTER TABLE events
DROP COLUMN IF EXISTS active,
DROP COLUMN IF EXISTS deleted_at,
DROP COLUMN IF EXISTS updated_at;
