-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT DEFAULT 'user', -- Roles like 'admin', 'moderator', or 'user'
    status TEXT DEFAULT 'offline', -- User online status: 'online', 'offline', 'away'
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    room_id INT NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    sent_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE TABLE user_room (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    room_id INT NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    joined_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, room_id)
);
-- +goose Down
DROP TABLE IF EXISTS user_room;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS rooms;
DROP TABLE IF EXISTS users;



