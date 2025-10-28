CREATE TABLE project_posts (
    Id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    photo_url TEXT,
    like_count integer DEFAULT 0 NOT NULL, 
    project_id INTEGER NOT NULL REFERENCES projects(Id) ON DELETE CASCADE
);

CREATE TABLE comments (
    Id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    like_count INTEGER DEFAULT 0 NOT NULL,
    post_id INTEGER NOT NULL REFERENCES project_posts(Id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(Id) ON DELETE CASCADE
);