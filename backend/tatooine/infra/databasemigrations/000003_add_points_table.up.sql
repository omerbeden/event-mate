CREATE TABLE IF NOT EXISTS user_points (
    id SERIAL PRIMARY KEY,
    giver_id INT NOT NULL,
    receiver_id INT NOT NULL,
    points NUMERIC(3,1) NOT NULL,
    comment TEXT,
    related_activity_id INT,
    given_at DATE DEFAULT NOW(), 
    UNIQUE(giver_id, receiver_id),
    FOREIGN KEY (giver_id) REFERENCES user_profiles(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_id) REFERENCES user_profiles(id) ON DELETE CASCADE,
    FOREIGN KEY (related_activity_id) REFERENCES activities(id) ON DELETE CASCADE
);
