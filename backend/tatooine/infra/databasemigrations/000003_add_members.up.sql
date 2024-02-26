ALTER TABLE activities
 ADD COLUMN content TEXT,
 ADD COLUMN start_at TIMESTAMP,
 ADD COLUMN end_at TIMESTAMP;


ALTER TABLE user_profiles
 ADD COLUMN profile_image_url TEXT