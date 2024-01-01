ALTER TABLE activities
 ADD COLUMN background_image_url TEXT,
 ADD COLUMN start_at TIMESTAMP,
 ADD COLUMN content TEXT;

ALTER TABLE user_profiles
 ADD COLUMN profile_image_url TEXT,
 ADD COLUMN profile_point INTEGER;
