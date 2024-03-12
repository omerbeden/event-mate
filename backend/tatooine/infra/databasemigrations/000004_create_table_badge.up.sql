CREATE TABLE IF NOT EXISTS profile_badges(
    id SERIAL PRIMARY KEY,
    profile_id INT NOT NULL,
    badge_id INT NOT NULL,
    image_url Text NOT NULL,
    text Text NOT NULL,
    given_at DATE DEFAULT NOW(),
    FOREIGN KEY (profile_id) REFERENCES user_profiles(id) ON DELETE CASCADE
)







