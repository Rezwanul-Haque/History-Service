CREATE DATABASE IF NOT EXISTS hds_db;

USE hds_db;

-- SET sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

CREATE TABLE `location_history` (
                                  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT 'Primary key.',
                                  `domain` varchar(255) NOT NULL COMMENT 'FK to the company.',
                                  `user_id` varchar(255) NOT NULL COMMENT 'User Id.',
                                  `client_timestamp_utc` timestamp NOT NULL COMMENT 'Client timestamp of when location data was generated.',
                                  `server_timestamp_utc` timestamp NOT NULL COMMENT 'Server timestamp of when location data was received.',
                                  `longitude` double DEFAULT NULL COMMENT 'Longitude of the location.',
                                  `latitude` double DEFAULT NULL COMMENT 'Latitude of the location.',
                                  PRIMARY KEY (`id`),
                                  KEY `idx_location_history_domain` (`domain`),
                                  KEY `idx_location_history_domain_user_id` (`domain`,`user_id`),
                                  CONSTRAINT `fk_location_history_domain` FOREIGN KEY (`domain`) REFERENCES `ids_db`.`company` (`domain`)
) ENGINE=InnoDB AUTO_INCREMENT=123 DEFAULT CHARSET=utf8;

CREATE USER 'hd_user'@'%' IDENTIFIED BY '12345678';
GRANT ALL PRIVILEGES ON *.* TO 'hd_user'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
