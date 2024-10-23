-- create database and tables
CREATE DATABASE IF NOT EXISTS airline_checkin;

USE airline_checkin;

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS seats (
    seat_no VARCHAR(3) PRIMARY KEY,
    user_id INT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);


INSERT INTO users (name) VALUES
('Sierra Kassulke'), ('Everette Douglas'), ('Quinn Kohler'), ('Theresa Sporer'), ('Hollie Gulgowski'), ('Ferne Hudson'), ('Jo Windler'), ('Helmer Ferry'), ('Wilma Waelchi'), ('Hayden Altenwerth'), ('Isaias Schiller'), ('Ansel Barrows'), ('Roy McCullough'), ('Alexys Bernhard'), ('Kristopher Beier'), ('Mason Rutherford'), ('Raul Reynolds'), ('Enrico Beer'), ('Kathleen Howe'), ('Morgan Kautzer'), ('Russel Kertzmann'), ('Magnus Casper'), ('Marina Pfannerstill'), ('Gordon Sawayn'), ('Berneice Witting'), ('Jacinto Schimmel'), ('Herminia Hansen'), ('Verner Mann'), ('Lauretta Krajcik'), ('Rafael Kihn'), ('Jeromy Mante'), ('Melyssa Nader'), ('Roma Schowalter'), ('Austen Mayert'), ('Bernard Schmeler'), ('Alfreda Orn'), ('Ryan Shanahan'), ('Natalia Jewess'), ('Dorcas Raynor'), ('Princess Lind'), ('Jessyca Schimmel'), ('Mozelle Kertzmann'), ('Rahul Willms'), ('Clementine Torp'), ('Sim Williamson'), ('Hermann Parisian'), ('Xavier Roob'), ('Cierra Hammes'), ('Mariam Adams'), ('Cara Heidenreich'), ('Noelia Bernhard'), ('Daren McDermott'), ('Tamia Stracke'), ('Rusty Ebert'), ('Anya Torp'), ('Vella Olson'), ('Raymond Daugherty'), ('Jadon Weissnat'), ('Wilmer Kutch'), ('Alex Heaney'), ('Mckenna Kulas'), ('Addison Auer'), ('Leonora Pacocha'), ('Breanna Schoen'), ('Jazlyn Windler'), ('Sidney Pollich'), ('Jakayla Zulauf'), ('Tessie Spencer'), ('Phyllis Romaguera'), ('Breana Jenkins'), ('Jayme Champlin'), ('Lorena Rosenbaum'), ('Jarrod Mosciski'), ('Irwin Rogahn'), ('Ludwig McClure'), ('Reymundo McDermott'), ('Dortha Tillman'), ('Jerel Stroman'), ('Barrett Quitzon'), ('Madonna Spencer'), ('Delfina Hegmann'), ('Leo Bergstrom'), ('Destany Conn'), ('Johathan Williamson'), ('Vince Donnelly'), ('Telly Kihn'), ('Cecil Muller'), ('Kale Turcotte'), ('Devin Schamberger'), ('Ali Schaden'), ('Israel Gorczany'), ('Wilfred Bayer'), ('Mac Jast'), ('Grant Jast'), ('Eugenia Olson'), ('Maximus Huel'), ('Hortense Volkman'), ('Isobel Kohler'), ('Alivia Gusikowski'), ('Deborah Gleichner'), ('Ashtyn Weber'), ('Jaiden Beatty'), ('Minnie Durgan'), ('Maybell Batz'), ('Jose Frami'), ('Sheila Mann'), ('Jaeden Block'), ('Kali Heaney'), ('Haskell Gottlieb'), ('Gerda Waelchi'), ('Triston Bashirian'), ('Bennett Hickle'), ('Nettie Heathcote'), ('Xander Ledner'), ('Taya Rosenbaum'), ('Kiana Prosacco'), ('Oswaldo Cole'), ('Euna Flatley'), ('Carlo Trantow'), ('Jammie Nienow');

INSERT INTO seats (seat_no) VALUES
('1A'), ('1B'), ('1C'), ('1D'), ('1E'), ('1F'), ('2A'), ('2B'), ('2C'), ('2D'), ('2E'), ('2F'), ('3A'), ('3B'), ('3C'), ('3D'), ('3E'), ('3F'), ('4A'), ('4B'), ('4C'), ('4D'), ('4E'), ('4F'), ('5A'), ('5B'), ('5C'), ('5D'), ('5E'), ('5F'), ('6A'), ('6B'), ('6C'), ('6D'), ('6E'), ('6F'), ('7A'), ('7B'), ('7C'), ('7D'), ('7E'), ('7F'), ('8A'), ('8B'), ('8C'), ('8D'), ('8E'), ('8F'), ('9A'), ('9B'), ('9C'), ('9D'), ('9E'), ('9F'), ('10A'), ('10B'), ('10C'), ('10D'), ('10E'), ('10F'), ('11A'), ('11B'), ('11C'), ('11D'), ('11E'), ('11F'), ('12A'), ('12B'), ('12C'), ('12D'), ('12E'), ('12F'), ('13A'), ('13B'), ('13C'), ('13D'), ('13E'), ('13F'), ('14A'), ('14B'), ('14C'), ('14D'), ('14E'), ('14F'), ('15A'), ('15B'), ('15C'), ('15D'), ('15E'), ('15F'), ('16A'), ('16B'), ('16C'), ('16D'), ('16E'), ('16F'), ('17A'), ('17B'), ('17C'), ('17D'), ('17E'), ('17F'), ('18A'), ('18B'), ('18C'), ('18D'), ('18E'), ('18F'), ('19A'), ('19B'), ('19C'), ('19D'), ('19E'), ('19F'), ('20A'), ('20B'), ('20C'), ('20D'), ('20E'), ('20F');
