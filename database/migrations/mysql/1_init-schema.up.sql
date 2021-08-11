CREATE TABLE IF NOT EXISTS users (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    `uuid` CHAR(36) UNIQUE,
    email VARCHAR(64),
    `password` VARCHAR(128),
    `name` VARCHAR(64),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);