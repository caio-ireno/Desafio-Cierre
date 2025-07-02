DROP DATABASE IF EXISTS mysqlDB;
CREATE DATABASE mysqlDB CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE mysqlDB;

DROP TABLE IF EXISTS `tickets`;

CREATE TABLE tickets(
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`name` varchar(100) NOT NULL,
`email` varchar(100) NOT NULL,
`country` varchar(100) NOT NULL,
`hour` varchar(100) NOT NULL,
`price` int(10) NOT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

