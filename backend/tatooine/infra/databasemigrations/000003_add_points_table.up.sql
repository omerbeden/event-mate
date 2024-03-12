CREATE TABLE IF NOT EXISTS user_points (
    id SERIAL PRIMARY KEY,
    giver_id varchar NOT NULL,
    receiver_id varchar NOT NULL,
    points NUMERIC(3,1) NOT NULL,
    comment TEXT,
    related_activity_id INT,
    given_on DATE DEFAULT NOW(), 
    UNIQUE(giver_id, receiver_id),
    FOREIGN KEY (giver_id) REFERENCES user_profiles(external_id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_id) REFERENCES user_profiles(external_id) ON DELETE CASCADE,
    FOREIGN KEY (related_activity_id) REFERENCES activities(id) ON DELETE CASCADE
);
