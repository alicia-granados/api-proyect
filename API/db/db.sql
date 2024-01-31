-- Create the database if it does not exist
CREATE DATABASE IF NOT EXISTS game_champion;

-- Use the database
USE game_champion;

-- Create the Champion table
CREATE TABLE Champion (
    Id    INT AUTO_INCREMENT PRIMARY KEY,
    Name  VARCHAR(255),
    Title VARCHAR(255),
    Lore  TEXT
);

-- Create the Skins table
CREATE TABLE Skins (
    Id          INT AUTO_INCREMENT PRIMARY KEY,
    Id_Num      VARCHAR(255) UNIQUE,
    Num         INT,
    Id_Champion INT,
    Name        VARCHAR(255)
);

-- Create the Tags table
CREATE TABLE Tags (
    Id          INT AUTO_INCREMENT PRIMARY KEY,
    Id_Champion INT,
    Name        VARCHAR(255)
);