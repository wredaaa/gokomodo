# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.4.17-MariaDB)
# Database: gokomodo
# Generation Time: 2022-12-01 14:40:38 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table buyers
# ------------------------------------------------------------

DROP TABLE IF EXISTS `buyers`;

CREATE TABLE `buyers` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL DEFAULT '',
  `name` varchar(100) NOT NULL DEFAULT '',
  `password` varchar(100) NOT NULL DEFAULT '',
  `shipping_address` text DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `buyers` WRITE;
/*!40000 ALTER TABLE `buyers` DISABLE KEYS */;

INSERT INTO `buyers` (`id`, `email`, `name`, `password`, `shipping_address`, `created_at`, `updated_at`)
VALUES
	(1,'wredaaa@gmail.com','Wreda Prasetiyo W','$2a$14$4OPJ1Qp7wPfVpbHY/WPjo.I2c/KuwrYYqxG9M9Y6sVC92KKBic/aG','Lumajang',NULL,NULL);

/*!40000 ALTER TABLE `buyers` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table orders
# ------------------------------------------------------------

DROP TABLE IF EXISTS `orders`;

CREATE TABLE `orders` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `buyer_id` int(11) DEFAULT NULL,
  `seller_id` int(11) DEFAULT NULL,
  `product_id` int(11) DEFAULT NULL,
  `delivery_source` text DEFAULT NULL,
  `delivery_destination` text DEFAULT NULL,
  `items` varchar(50) DEFAULT NULL,
  `price` decimal(11,0) DEFAULT NULL,
  `quantity` int(11) DEFAULT NULL,
  `total_price` decimal(11,0) DEFAULT NULL,
  `status` varchar(10) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;

INSERT INTO `orders` (`id`, `buyer_id`, `seller_id`, `product_id`, `delivery_source`, `delivery_destination`, `items`, `price`, `quantity`, `total_price`, `status`, `created_at`, `updated_at`)
VALUES
	(1,1,1,1,'','','pisang',10000,5,50000,'Accepted','2022-11-30 20:01:37','2022-12-01 00:17:07'),
	(2,1,1,1,'jakarta','yogya','pisang',10000,5,50000,'Pending','2022-11-30 20:02:10','2022-11-30 20:02:10');

/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table products
# ------------------------------------------------------------

DROP TABLE IF EXISTS `products`;

CREATE TABLE `products` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `seller_id` int(11) NOT NULL,
  `product_name` varchar(50) NOT NULL DEFAULT '',
  `description` text DEFAULT NULL,
  `price` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;

INSERT INTO `products` (`id`, `seller_id`, `product_name`, `description`, `price`, `created_at`, `updated_at`)
VALUES
	(1,1,'pisang','asli lumajang 100%',10000,'2022-11-30 17:55:08','2022-11-30 18:01:31'),
	(2,1,'pisang','asli lumajang',10000,'2022-11-30 17:55:49','2022-11-30 17:55:49'),
	(3,0,'','',0,'2022-12-01 01:40:06','2022-12-01 01:40:06');

/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sellers
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sellers`;

CREATE TABLE `sellers` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL DEFAULT '',
  `name` varchar(100) NOT NULL DEFAULT '',
  `password` varchar(100) NOT NULL DEFAULT '',
  `pickup_address` text DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `sellers` WRITE;
/*!40000 ALTER TABLE `sellers` DISABLE KEYS */;

INSERT INTO `sellers` (`id`, `email`, `name`, `password`, `pickup_address`, `created_at`, `updated_at`)
VALUES
	(1,'yusuf@gmail.com','yusuf','$2a$14$4OPJ1Qp7wPfVpbHY/WPjo.I2c/KuwrYYqxG9M9Y6sVC92KKBic/aG',NULL,NULL,NULL);

/*!40000 ALTER TABLE `sellers` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
