CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    from_user_id INT REFERENCES users(id),
    to_user_id INT REFERENCES users(id),
    amount INT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
