ALTER TABLE user_profiles
DROP COLUMN IF EXISTS external_id,
DROP COLUMN IF EXISTS user_name;