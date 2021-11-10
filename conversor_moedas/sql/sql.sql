CREATE DATABASE IF NOT EXISTS conversor_moedas;

USE conversor_moedas;

DROP TABLE IF EXISTS depositos;

CREATE TABLE depositos (
    id int auto_increment primary key,
    ValorDepositado float not null,
    criadoEm timestamp default current_timestamp() not null
)ENGINE=INNODB;