-- +goose Up
ALTER TABLE users ADD COLUMN api_key TEXT DEFAULT NULL;
ALTER TABLE users ADD CONSTRAINT unique_api_key UNIQUE (api_key);
-- +goose Down
ALTER TABLE users DROP CONSTRAINT unique_api_key;
ALTER TABLE users DROP COLUMN api_key;
