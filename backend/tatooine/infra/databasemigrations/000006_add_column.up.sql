ALTER TABLE user_profiles
ADD COLUMN external_id Text UNIQUE,
ADD COLUMN user_name varchar(20) UNIQUE;