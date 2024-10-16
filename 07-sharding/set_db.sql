CREATE DATABASE IF NOT EXISTS test_db;

USE test_db;

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO users (id, name) VALUES (1, 'John Doe');
INSERT INTO users (id, name) VALUES (2, 'Jane Smith');
INSERT INTO users (id, name) VALUES (3, 'Dr. Strange');
INSERT INTO users (id, name) VALUES (4, 'Octo Octavious');
INSERT INTO users (id, name) VALUES (5, 'Pepper Pots');
INSERT INTO users (id, name) VALUES (6, 'Sheldon Copper');
INSERT INTO users (id, name) VALUES (7, 'Viva La Da');
INSERT INTO users (id, name) VALUES (8, 'Peter Paker');
