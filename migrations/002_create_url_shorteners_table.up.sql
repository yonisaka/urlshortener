CREATE TABLE url_shorteners (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    original_url VARCHAR(255) NOT NULL,
    shortened_url VARCHAR(255) NOT NULL,
    datetime TIMESTAMPTZ NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
