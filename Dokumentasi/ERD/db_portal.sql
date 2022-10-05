/*
SQLyog Ultimate
MySQL - 5.7.24 : Database - db_portal
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`db_portal` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `db_portal`;

/*Table structure for table `contents` */

DROP TABLE IF EXISTS `contents`;

CREATE TABLE `contents` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext,
  `content` longtext,
  `images` longtext,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_contents_deleted_at` (`deleted_at`),
  KEY `fk_contents_user` (`user_id`),
  CONSTRAINT `fk_contents_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

/*Data for the table `contents` */

insert  into `contents`(`id`,`created_at`,`updated_at`,`deleted_at`,`title`,`content`,`images`,`user_id`) values 
(1,'2022-10-05 15:08:12.817','2022-10-05 15:08:12.817',NULL,'berita hangat','berita terkini kedua','https://www.kindpng.com/picc/m/151-1517817_transparent-computadoras-png-desktop-computer-png-download.png',3),
(2,'2022-10-05 15:12:12.869','2022-10-05 17:36:25.198',NULL,'ini judul','berita terkini keduakali','https://www.kindpng.com/picc/m/151-1517817_transparent-computadoras-png-desktop-computer-png-download.png',3);

/*Table structure for table `details` */

DROP TABLE IF EXISTS `details`;

CREATE TABLE `details` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `point` bigint(20) DEFAULT NULL,
  `views_number` bigint(20) DEFAULT NULL,
  `content_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_details_deleted_at` (`deleted_at`),
  KEY `fk_details_content` (`content_id`),
  CONSTRAINT `fk_details_content` FOREIGN KEY (`content_id`) REFERENCES `contents` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

/*Data for the table `details` */

insert  into `details`(`id`,`created_at`,`updated_at`,`deleted_at`,`point`,`views_number`,`content_id`) values 
(3,'2022-10-05 15:52:17.137','2022-10-05 15:52:17.137',NULL,111,600,2),
(4,'2022-10-05 15:53:10.990','2022-10-05 15:53:10.990',NULL,111,600,1);

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` longtext,
  `email` varchar(191) DEFAULT NULL,
  `password` longtext,
  `full_name` longtext,
  `role` varchar(191) DEFAULT 'creator',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

/*Data for the table `users` */

insert  into `users`(`id`,`created_at`,`updated_at`,`deleted_at`,`username`,`email`,`password`,`full_name`,`role`) values 
(2,'2022-10-01 07:41:20.261','2022-10-01 07:41:20.261',NULL,'administrator','admin@gmail.com','$2a$10$OCBVBqv8NCrKEt3wMIHrnemgz90kAZEtV.XhjthAL5R0Q2YkgloRa','Admin','admin'),
(3,'2022-10-01 07:43:40.785','2022-10-01 07:43:40.785',NULL,'reski','programmer.reski@gmail.com','$2a$10$iWy2iINk.F.RZZ2XIMgO5umawpQCuaiauGM5Yi17/83C6lvOy2BL2','Ahmad Reski','creator'),
(4,'2022-10-05 17:41:40.399','2022-10-05 17:41:40.399',NULL,'batman','batman.reski@gmail.com','$2a$10$vNXQ3vgxg.6NWWLPelPClOjatGEwAnWZh.0FOLxq8NJ4Q37r.cuiK','batman','creator');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
