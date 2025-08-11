-- simple migration for users table
CREATE TABLE IF NOT EXISTS users (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL,
email TEXT NOT NULL UNIQUE,
created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);