CREATE TABLE purchases (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    item_name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    purchased_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
