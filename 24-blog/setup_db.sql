CREATE DATABASE blog_shard1;
CREATE DATABASE blog_shard2;

USE blog_shard1;
CREATE TABLE blogs (
    blog_id BIGINT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Repeat for blog_shard2, or more shards as needed.
USE blog_shard2;
CREATE TABLE blogs (
    blog_id BIGINT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);