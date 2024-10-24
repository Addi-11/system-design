CREATE DATABASE IF NOT EXISTS test_db;

USE test_db;

-- Create table for shard 1
CREATE TABLE kv_store_shard1 (
  `key` VARCHAR(255) PRIMARY KEY,
  `value` TEXT,
  `expired_at` DATETIME
);

-- Create table for shard 2
CREATE TABLE kv_store_shard2 (
  `key` VARCHAR(255) PRIMARY KEY,
  `value` TEXT,
  `expired_at` DATETIME
);
