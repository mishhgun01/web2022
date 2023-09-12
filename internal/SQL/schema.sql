CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    login TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT '',
    last_path TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS notes(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    description TEXT,
    status INTEGER NOT NULL DEFAULT 0,
    user_id INTEGER NOT NULL references users(id)
);