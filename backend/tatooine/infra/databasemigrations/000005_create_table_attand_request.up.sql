CREATE TABLE IF NOT EXISTS attend_requests (
    id SERIAL PRIMARY KEY,
    sender_id INT REFERENCES user_profiles(id),
    receiver_id INT REFERENCES user_profiles(id),
    status VARCHAR(20) DEFAULT 'pending',
    created_at DATE DEFAULT CURRENT_TIMESTAMP
);