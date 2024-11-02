CREATE DATABASE IF NOT EXISTS zomato;

USE zomato;

-- NOTE: Adding auto-increment might be better - to ids

-- keeps track of delivery agents
CREATE TABLE IF NOT EXISTS agents (
    id INT PRIMARY KEY,
    is_reserved BOOLEAN DEFAULT FALSE,
    order_id VARCHAR(100),
    name VARCHAR(30)
);

-- keep tracks of food items avaible, that the user can order
CREATE TABLE IF NOT EXISTS foods (
    id INT PRIMARY KEY,
    name VARCHAR(50)
);

-- keeps track of the food delivery packets
CREATE TABLE IF NOT EXISTS packets (
    id INT PRIMARY KEY,
    food_id INT,
    is_reserved BOOLEAN DEFAULT FALSE,
    order_id VARCHAR(100),
    FOREIGN KEY (food_id) REFERENCES foods(id)
);

-- entering food items
INSERT INTO foods VALUES (1, 'burger');
INSERT INTO foods VALUES (2, 'pizza');

-- inserting into agents
INSERT INTO agents (id, name) VALUES
    (1, 'Agent A'),
    (2, 'Agent B'),
    (3, 'Agent C'),
    (4, 'Agent D'),
    (5, 'Agent E'),
    (6, 'Agent F'),
    (7, 'Agent G'),
    (8, 'Agent H'),
    (9, 'Agent I'),
    (10, 'Agent J');

-- inserting into packets different food items user wants
INSERT INTO packets (id, food_id) VALUES
    (1, 1),
    (2, 2),
    (3, 1),
    (4, 2),
    (5, 1),
    (6, 2),
    (7, 1),
    (8, 2),
    (9, 1),
    (10, 2),
    (11, 1),
    (12, 2),
    (13, 1),
    (14, 2),
    (15, 1),
    (16, 2),
    (17, 1),
    (18, 2),
    (19, 1),
    (20, 2);
