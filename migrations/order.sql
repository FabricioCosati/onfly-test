CREATE DATABASE IF NOT EXISTS onflydb;
USE onflydb;

CREATE TABLE `onflydb`.`orders` (
  `id_order` INT NOT NULL AUTO_INCREMENT,
  `requester_name_order` VARCHAR(255) NOT NULL,
  `destination_order` VARCHAR(255) NOT NULL,
  `going_date_order` DATE NOT NULL,
  `return_date_order` DATE NOT NULL,
  `status_order` ENUM("requested", "approved", "canceled") NOT NULL,
  `created_at_order` DATETIME NOT NULL DEFAULT NOW(),
  `updated_at_order` DATETIME NOT NULL DEFAULT NOW(),
  PRIMARY KEY (`id_order`));
