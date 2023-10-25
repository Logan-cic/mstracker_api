CREATE DATABASE IF NOT EXISTS mstracker;

USE mstracker;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(20) NOT NULL UNIQUE,
    CriadoEm TIMESTAMP DEFAULT current_timestamp()
) ENGINE=INNODB;