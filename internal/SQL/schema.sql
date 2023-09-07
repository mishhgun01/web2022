CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    login TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT '',
    session_id TEXT NOT NULL DEFAULT ''
);

CREATE TABLE notes IF NOT EXISTS(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    description TEXT,
    status INTEGER NOT NULL DEFAULT 0,
    user_id INTEGER NOT NULL references users(id)
);