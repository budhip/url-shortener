CREATE DATABASE  IF NOT EXISTS `shortener` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `shortener`;

DROP TABLE IF EXISTS `url_shortener`;

CREATE TABLE `url_shortener` (
	`id` int not null auto_increment,
	`slug` varchar(14) collate utf8mb4_unicode_ci NOT NULL,
	`url` varchar(620) collate utf8mb4_unicode_ci NOT NULL,
	`created_at` datetime NOT NULL,
	`last_seen` datetime NOT NULL default '1000-01-01 00:00:00',
	`hits` bigint(20) NOT NULL default 0,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;