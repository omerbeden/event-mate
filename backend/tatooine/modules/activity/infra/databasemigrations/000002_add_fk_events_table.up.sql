ALTER TABLE activities ADD CONSTRAINT fk_events_user_profiles FOREIGN KEY (created_user_id) REFERENCES user_profiles(id);