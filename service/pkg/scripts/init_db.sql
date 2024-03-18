CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


-- Create users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

-- Create movies table
CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    imglink TEXT,
    release_date DATE,
    summary TEXT,
    genre VARCHAR(255),
    rating INTEGER DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
