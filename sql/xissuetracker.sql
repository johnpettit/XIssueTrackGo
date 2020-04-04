CREATE DATABASE IF NOT EXISTS `XIssueTracker` /*!40100 DEFAULT CHARACTER SET utf8 */;

/* CREATE TABLES */

use `XIssueTracker`;

/* user Table */

CREATE TABLE IF NOT EXISTS `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `firstname` varchar(45) DEFAULT NULL,
  `lastname` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `password` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* issue Table */

CREATE TABLE IF NOT EXISTS `issue` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* token Table */
CREATE TABLE `token` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tokenhash` varchar(240) NOT NULL,
  `userid` int(10) unsigned NOT NULL,
  `lasttouch` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
