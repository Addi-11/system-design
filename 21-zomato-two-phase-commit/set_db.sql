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
INSERT INTO foods VALUES (3, 'tandoori chicken');
INSERT INTO foods VALUES (4, 'pasta');

-- inserting into agents
INSERT INTO agents (id, name) VALUES
    (1, 'Sachin'),
    (2, 'Monika'),
    (3, 'John'),
    (4, 'Stacy'),
    (5, 'Priya'),
    (6, 'Joe'),
    (7, 'Rachel'),
    (8, 'Ross'),
    (9, 'Chandler'),
    (10, 'Barney');

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
