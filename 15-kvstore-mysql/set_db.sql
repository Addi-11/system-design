CREATE DATABASE IF NOT EXISTS test_db;

USE test_db;

CREATE TABLE IF NOT EXISTS kv_store (
    `key` VARCHAR(255) PRIMARY KEY,
    `value` TEXT NOT NULL,
    `expired_at` DATETIME
);