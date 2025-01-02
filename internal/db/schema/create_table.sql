CREATE TABLE url_mapping (
                             id BIGINT PRIMARY KEY, -- Optional: A unique ID for each row
                             long_url TEXT NOT NULL UNIQUE, -- Column for storing long URLs, not null and unique
                             short_url VARCHAR(8) NOT NULL UNIQUE, -- Column for storing short URLs, not null and unique
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Column for storing the creation timestamp with default current time
);