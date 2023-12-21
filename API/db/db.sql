-- Crear la base de datos si no existe
CREATE DATABASE IF NOT EXISTS game_champion;

-- Utilizar la base de datos
USE game_champion;

-- Crear la tabla Champion
CREATE TABLE Champion (
    Id    INT AUTO_INCREMENT PRIMARY KEY,
    Name  VARCHAR(255),
    Title VARCHAR(255),
    Lore  TEXT
);

-- Crear la tabla Skins
CREATE TABLE Skins (
    Id   INT AUTO_INCREMENT PRIMARY KEY,
    Id_Num  VARCHAR(255) UNIQUE,
    Num    INT ,
    Id_Champion INT, 
    Name VARCHAR(255) 
);

-- Crear la tabla Tags
CREATE TABLE Tags (
    Id   INT AUTO_INCREMENT PRIMARY KEY,
    Id_Champion INT,
    Name VARCHAR(255) UNIQUE
);
