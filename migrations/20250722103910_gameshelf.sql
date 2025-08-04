-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'game_status') THEN
        CREATE TYPE game_status AS ENUM ('completed', 'in_progress', 'abandoned', 'not_owned');
    END IF;
END
$$;
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'game_store') THEN
        CREATE TYPE game_store AS ENUM ('steam', 'epic_store', 'xbox',
         'playstation', 'nintendo', 'microsoft_store');
    END IF;
END
$$;
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'friend_req_status') THEN
        CREATE TYPE friend_req_status AS ENUM ('pending', 'accepted', 'declined');
    END IF;
END
$$;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS games (
    title TEXT UNIQUE NOT NULL PRIMARY KEY,
    genre TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_games (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL REFERENCES users(username) ON DELETE CASCADE,
    game_title TEXT NOT NULL REFERENCES games(title) ON DELETE CASCADE,
    game_status game_status NOT NULL,
    game_store game_store NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(username, game_title)
);

CREATE TABLE IF NOT EXISTS friend_requests (
    id SERIAL PRIMARY KEY,
    sender TEXT NOT NULL REFERENCES users(username) ON DELETE CASCADE,
    receiver TEXT NOT NULL REFERENCES users(username) ON DELETE CASCADE,
    req_status friend_req_status NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (sender, receiver)
);

CREATE TABLE IF NOT EXISTS friends (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL REFERENCES users(username) ON DELETE CASCADE,
    friend TEXT NOT NULL REFERENCES users(username) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (username, friend)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_games;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS games;
DROP TYPE IF EXISTS game_status;
DROP TYPE IF EXISTS friend_requests;
DROP TYPE IF EXISTS friens;
-- +goose StatementEnd
