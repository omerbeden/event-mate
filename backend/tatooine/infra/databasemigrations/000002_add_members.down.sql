ALTER TABLE activities
DROP COLUMN IF EXISTS content,
DROP COLUMN IF EXISTS start_at;
DROP COLUMN IF EXISTS end_at;

