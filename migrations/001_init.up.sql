CREATE TYPE user_role AS ENUM ('admin', 'investor', 'founder');

CREATE TABLE users (
    Id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    role user_role NOT NULL DEFAULT 'investor',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE companies (
    Id SERIAL PRIMARY KEY,
    name TEXT NOT NULL  UNIQUE,
    user_id INTEGER NOT NULL REFERENCES users(Id) ON DELETE CASCADE
);

CREATE TABLE projects (
    Id SERIAL PRIMARY KEY,
    title TEXT NOT NULL UNIQUE,
    description TEXT,
    goal_amount NUMERIC(12,2) NOT NULL,
    collected_amount NUMERIC(12,2) DEFAULT 0,
    company_id INTEGER NOT NULL REFERENCES companies(Id) ON DELETE CASCADE
);