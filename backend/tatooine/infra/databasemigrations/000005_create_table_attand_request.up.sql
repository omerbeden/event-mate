CREATE TABLE attend_requests IF NOT EXISTS(
    id SERIAL PRIMARY KEY,
    sender_id INT REFERENCES user_profiles(id),
    receiver_id INT REFERENCES user_profiles(id),
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);