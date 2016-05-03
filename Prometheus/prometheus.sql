-- MySQL dump 10.13  Distrib 5.1.73, for redhat-linux-gnu (x86_64)
--
-- Host: 192.168.33.13    Database: prometheus
-- ------------------------------------------------------
-- Server version	5.6.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `cabinet`
--

DROP TABLE IF EXISTS `cabinet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cabinet` (
  `cabinet_id` int(11) NOT NULL AUTO_INCREMENT,
  `cabinet_name` varchar(255) NOT NULL,
  `iscloud` enum('true','false') NOT NULL,
  `capacity_total` bigint(20) DEFAULT NULL,
  `capacity_used` bigint(20) DEFAULT NULL,
  `location_id` int(11) NOT NULL,
  PRIMARY KEY (`cabinet_id`),
  UNIQUE KEY `uniq_name_location_id` (`cabinet_name`,`location_id`),
  KEY `catbinet_fk_location_id` (`location_id`),
  CONSTRAINT `catbinet_fk_location_id` FOREIGN KEY (`location_id`) REFERENCES `location` (`location_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cabinet`
--

LOCK TABLES `cabinet` WRITE;
/*!40000 ALTER TABLE `cabinet` DISABLE KEYS */;
INSERT INTO `cabinet` VALUES (1,'test-cabinet','false',50,0,1),(2,'test-cabinet2','false',50,0,1),(3,'test-cabinet3','false',50,0,2);
/*!40000 ALTER TABLE `cabinet` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `device`
--

DROP TABLE IF EXISTS `device`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `device` (
  `device_id` int(11) NOT NULL AUTO_INCREMENT,
  `device_name` varchar(255) DEFAULT NULL,
  `device_model_id` int(11) DEFAULT NULL,
  `father_device_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`device_id`),
  UNIQUE KEY `uniq_name` (`device_name`) USING BTREE,
  KEY `device_fk_father_device_id` (`father_device_id`),
  KEY `device_fk_device_model_id` (`device_model_id`) USING BTREE,
  CONSTRAINT `device_fk_device_model_id` FOREIGN KEY (`device_model_id`) REFERENCES `deviceModel` (`device_model_id`) ON UPDATE CASCADE,
  CONSTRAINT `device_fk_father_device_id` FOREIGN KEY (`father_device_id`) REFERENCES `device` (`device_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `device`
--

LOCK TABLES `device` WRITE;
/*!40000 ALTER TABLE `device` DISABLE KEYS */;
INSERT INTO `device` VALUES (1,'test-test192.168.33.81',1,NULL);
/*!40000 ALTER TABLE `device` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `deviceModel`
--

DROP TABLE IF EXISTS `deviceModel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `deviceModel` (
  `device_model_id` int(11) NOT NULL AUTO_INCREMENT,
  `device_model_name` varchar(255) NOT NULL,
  `device_type` enum('Other','Router','Switch','Server','VM') NOT NULL DEFAULT 'Other',
  PRIMARY KEY (`device_model_id`),
  UNIQUE KEY `uniq_device_model_name` (`device_model_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deviceModel`
--

LOCK TABLES `deviceModel` WRITE;
/*!40000 ALTER TABLE `deviceModel` DISABLE KEYS */;
INSERT INTO `deviceModel` VALUES (1,'IBM','Server'),(2,'华为','Router'),(3,'cisco','Switch'),(11,'1asdasdasda','Server'),(13,'333','Router');
/*!40000 ALTER TABLE `deviceModel` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `deviceTag`
--

DROP TABLE IF EXISTS `deviceTag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `deviceTag` (
  `device_id` int(11) NOT NULL,
  `tag_id` int(11) NOT NULL,
  UNIQUE KEY `u_relation` (`device_id`,`tag_id`),
  KEY `fk_deviceTag_tag_id` (`tag_id`),
  CONSTRAINT `fk_deviceTag_device_id` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_deviceTag_tag_id` FOREIGN KEY (`tag_id`) REFERENCES `tag` (`tag_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deviceTag`
--

LOCK TABLES `deviceTag` WRITE;
/*!40000 ALTER TABLE `deviceTag` DISABLE KEYS */;
/*!40000 ALTER TABLE `deviceTag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `location`
--

DROP TABLE IF EXISTS `location`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `location` (
  `location_id` int(11) NOT NULL AUTO_INCREMENT,
  `location_name` varchar(255) NOT NULL,
  `picture` varchar(255) NOT NULL,
  `father_location_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`location_id`),
  UNIQUE KEY `uniq_name` (`location_name`) USING BTREE,
  KEY `fk_location_father_location_id` (`father_location_id`),
  CONSTRAINT `fk_location_father_location_id` FOREIGN KEY (`father_location_id`) REFERENCES `location` (`location_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `location`
--

LOCK TABLES `location` WRITE;
/*!40000 ALTER TABLE `location` DISABLE KEYS */;
INSERT INTO `location` VALUES (1,'大连IDC','1.jpg',2),(2,'长春IDC','1.jpg',NULL);
/*!40000 ALTER TABLE `location` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `netPort`
--

DROP TABLE IF EXISTS `netPort`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `netPort` (
  `netPort_id` int(11) NOT NULL,
  `mac` varchar(17) DEFAULT NULL,
  `ipv4_int` bigint(64) DEFAULT NULL,
  `device_id` int(11) DEFAULT NULL,
  `netPort_type` enum('vip','eth','bond') NOT NULL,
  UNIQUE KEY `un_ipv4_int` (`ipv4_int`),
  KEY `ipv4_fk_device_id` (`device_id`),
  CONSTRAINT `fk_netPort_device_id` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `netPort`
--

LOCK TABLES `netPort` WRITE;
/*!40000 ALTER TABLE `netPort` DISABLE KEYS */;
INSERT INTO `netPort` VALUES (0,'00:CF:E0:31:D5:D5',3232244049,1,'eth'),(1,'00:CF:E0:31:D5:D5',3232244050,1,'eth');
/*!40000 ALTER TABLE `netPort` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `space`
--

DROP TABLE IF EXISTS `space`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `space` (
  `cabinet_id` int(11) NOT NULL,
  `device_id` int(11) NOT NULL,
  `u_position` int(11) NOT NULL,
  `position` enum('rear','front','interior') NOT NULL,
  UNIQUE KEY `atom_fk_position` (`cabinet_id`,`u_position`,`position`),
  KEY `atom_fk_cabinet_id` (`cabinet_id`),
  KEY `atom_fk_device_id` (`device_id`),
  CONSTRAINT `atom_fk_cabinet_id` FOREIGN KEY (`cabinet_id`) REFERENCES `cabinet` (`cabinet_id`) ON UPDATE CASCADE,
  CONSTRAINT `atom_fk_device_id` FOREIGN KEY (`device_id`) REFERENCES `device` (`device_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `space`
--

LOCK TABLES `space` WRITE;
/*!40000 ALTER TABLE `space` DISABLE KEYS */;
/*!40000 ALTER TABLE `space` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag`
--

DROP TABLE IF EXISTS `tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tag` (
  `tag_id` int(11) NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(255) NOT NULL,
  PRIMARY KEY (`tag_id`),
  UNIQUE KEY `uniq_name` (`tag_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag`
--

LOCK TABLES `tag` WRITE;
/*!40000 ALTER TABLE `tag` DISABLE KEYS */;
INSERT INTO `tag` VALUES (1,'nginx');
/*!40000 ALTER TABLE `tag` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-05-03 17:39:24
